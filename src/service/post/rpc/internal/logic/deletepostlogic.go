package logic

import (
	"context"
	"encoding/json"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	po, err := l.svcCtx.PostInterface.GetPost(l.ctx, in.PostID)
	if err != nil {
		if po == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get post, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}
	// check permission
	if po.AuthorID != in.WxUid {
		return nil, errs.RpcPermissionDenied
	}
	poProj, err := l.svcCtx.PostInterface.DeletePost(l.ctx, in.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			//
			return nil, errs.RpcAlreadyExists
		}
		l.Logger.Errorf("cannot delete post, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	// send event to mq
	evt := event.DeletePost{
		Post:    poProj.Post,
		Project: poProj.Project,
	}
	bytes, _ := json.Marshal(&evt)

	err = l.svcCtx.KqDeletePostPusher.Push(l.ctx, string(bytes))
	if err != nil {
		// push to local queue
		topic := l.svcCtx.Config.KqDeletePostPusherConf.Topic
		l.Logger.
			WithFields(logx.Field("topic", topic)).
			Errorf("cannot push a message to mq when deleting post, err: %+V", err)

		l.svcCtx.LocalQueue.Push(evt)
	}

	return &post.DeletePostResponse{}, nil
}
