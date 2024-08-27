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

type GetPostsByAuthorIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsByAuthorIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsByAuthorIDLogic {
	return &GetPostsByAuthorIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsByAuthorIDLogic) GetPostsByAuthorID(in *post.GetPostsByAuthorIDRequest) (*post.GetPostsByAuthorIDResponse, error) {
	posts, paginator, err := l.svcCtx.PostInterface.GetPostsByAuthorID(l.ctx, in.Page, in.Size, in.AuthorID, params.ToOrderByOpt(in.OrderBy))
	if err != nil {
		log.Error().Msgf("cannot get posts by author_id, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfos := make([]*post.PostInfo, 0, len(posts))
	for _, po := range posts {
		poInfo := converter.PostDBToRPC(po)
		poInfos = append(poInfos, poInfo)
	}

	return &post.GetPostsByAuthorIDResponse{Posts: poInfos, Paginator: paginator.ToRPC()}, nil
}
