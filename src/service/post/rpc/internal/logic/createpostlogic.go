package logic

import (
	"context"
	"github.com/rs/zerolog/log"

	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/pb/post"
	"lookingforpartner/pkg/nanoid"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(in *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	var po *model.Post
	var proj model.Project
	var err error

	po = &model.Post{
		PostID:   constant.NanoidPrefixPost + nanoid.Gen(),
		AuthorID: in.WxUid,
		Title:    in.Title,
		Content:  in.Content,
	}

	if in.Project != nil {
		proj = model.Project{
			ProjectID:     constant.NanoidPrefixProject + nanoid.Gen(),
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
		log.Error().Msgf("cannot create post, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfo := converter.PostDBToRPC(po)

	return &post.CreatePostResponse{PostInfo: poInfo}, nil
}
