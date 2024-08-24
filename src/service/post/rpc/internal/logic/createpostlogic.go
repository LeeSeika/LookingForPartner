package logic

import (
	"context"

	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
	var po *model.Post
	var proj model.Project
	var err error

	po = &model.Post{
		PostID:   "1", //todo
		AuthorID: in.WxUid,
		Title:    in.Title,
		Content:  in.Content,
	}

	if in.Project != nil {
		proj = model.Project{
			ProjectID:     "1", //todo
			MaintainerID:  in.Project.Maintainer.WxUid,
			Name:          in.Project.Name,
			Introduction:  in.Project.Introduction,
			Role:          in.Project.Role,
			HeadCountInfo: in.Project.HeadCountInfo,
			Progress:      in.Project.Progress,
		}
		po.Project = proj
	}

	po, err = l.svcCtx.PostInterface.CreatePost(l.ctx, po)
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] CreatePostWithProjectTx error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfo := converter.PostDBToRPC(po)

	return &post.CreatePostResponse{PostInfo: poInfo}, nil
}
