package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/logic"
	"lookingforpartner/service/comment/rpc/internal/svc"
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
	if err := json.Unmarshal([]byte(val), &evt); err != nil {
		c.logger.
			Errorf("cannot unmarshal json, err: %+v", err)
		return err
	}

	createSubjectReq := comment.CreateSubjectRequest{
		PostID:         evt.Post.PostID,
		IdempotencyKey: evt.IdempotencyKey,
	}
	l := logic.NewCreateSubjectLogic(ctx, c.svcCtx)

	_, err := l.CreateSubject(&createSubjectReq)
	if err != nil {
		c.logger.Errorf("cannot create subject, err: %+v", err)
		return err
	}

	return nil
}
