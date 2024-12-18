package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/config"
	"lookingforpartner/service/user/rpc/internal/logic"
	"lookingforpartner/service/user/rpc/internal/svc"
)

type CreatePost struct {
	svcCtx *svc.ServiceContext
	conf   config.Config
	logger logx.Logger
}

func NewCreatePost(conf config.Config, svcCtx *svc.ServiceContext) *CreatePost {
	return &CreatePost{
		svcCtx: svcCtx,
		conf:   conf,
		logger: logx.WithContext(context.Background()).WithFields(logx.Field("consumer-topic", conf.KqCreatePostConsumerConf.Topic)),
	}
}

func (c *CreatePost) Consume(ctx context.Context, key, val string) error {

	var evt event.CreatePost
	err := json.Unmarshal([]byte(val), &evt)
	if err != nil {
		c.logger.
			Errorf("cannot unmarshal json, err: %+v", err)
		return err
	}

	updateUserPostCountReq := user.UpdateUserPostCountRequest{
		IdempotencyKey: evt.IdempotencyKey,
		WxUid:          evt.Post.AuthorID,
		Delta:          1,
	}

	l := logic.NewUpdateUserPostCountLogic(ctx, c.svcCtx)

	_, err = l.UpdateUserPostCount(&updateUserPostCountReq)
	if err != nil {
		c.logger.Errorf("cannot update user post count, err: %+v", err)
		return err
	}

	return nil
}
