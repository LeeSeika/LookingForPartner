package logic

import (
	"context"
	"github.com/rs/zerolog/log"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/params"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type GetPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsLogic) GetPosts(in *post.GetPostsRequest) (*post.GetPostsResponse, error) {
	posts, paginator, err := l.svcCtx.PostInterface.GetPosts(l.ctx, in.Page, in.Size, params.ToOrderByOpt(in.OrderBy))
	if err != nil {
		log.Error().Msgf("cannot get posts, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfos := make([]*post.PostInfo, 0, len(posts))
	for _, po := range posts {
		poInfo := converter.PostDBToRPC(po)
		poInfos = append(poInfos, poInfo)
	}

	return &post.GetPostsResponse{Posts: poInfos, Paginator: paginator.ToRPC()}, nil
}
