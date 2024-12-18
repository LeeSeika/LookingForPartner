package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/post"
	"lookingforpartner/pkg/nanoid"
	"lookingforpartner/service/post/model/entity"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	var po *entity.Post
	var proj *entity.Project
	var err error

	po = &entity.Post{
		PostID:   constant.NanoidPrefixPost + nanoid.Gen(),
		AuthorID: in.WxUid,
		Title:    in.Title,
		Content:  in.Content,
	}

	if in.Project != nil {
		proj = &entity.Project{
			ProjectID:     constant.NanoidPrefixProject + nanoid.Gen(),
			MaintainerID:  in.Project.Maintainer.WxUid,
			Name:          in.Project.Name,
			Introduction:  in.Project.Introduction,
			Role:          in.Project.Role,
			HeadCountInfo: in.Project.HeadCountInfo,
			Progress:      in.Project.Progress,
		}
	}

	poProj, err := l.svcCtx.PostInterface.CreatePost(l.ctx, po, proj, in.IdempotencyKey)
	if err != nil {
		l.Logger.Errorf("cannot create post, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	poInfo := converter.PostDBToRPC(poProj.Post)
	if poProj.Project != nil {
		projInfo := converter.ProjectDBToRPC(poProj.Project)
		poInfo.Project = projInfo
	}

	// send event to mq
	evt := event.CreatePost{
		IdempotencyKey: in.IdempotencyKey,
		Post:           po,
	}
	bytes, _ := json.Marshal(&evt)

	err = l.svcCtx.KqCreatePostPusher.Push(l.ctx, string(bytes))
	if err != nil {
		// push to local queue
		topic := l.svcCtx.Config.KqCreatePostPusherConf.Topic
		l.Logger.
			WithFields(logx.Field("topic", topic)).
			WithFields(logx.Field("idempotencyKey", in.IdempotencyKey)).
			Errorf("cannot push a message to mq when creating post, err: %+v", err)

		l.svcCtx.LocalQueue.Push(evt)
	}

	return &post.CreatePostResponse{PostInfo: poInfo}, nil
}
