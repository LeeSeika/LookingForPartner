package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/params"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsByAuthorIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostsByAuthorIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsByAuthorIDLogic {
	return &GetPostsByAuthorIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostsByAuthorIDLogic) GetPostsByAuthorID(in *post.GetPostsByAuthorIDRequest) (*post.GetPostsByAuthorIDResponse, error) {
	posts, paginator, err := l.svcCtx.PostInterface.GetPostsByAuthorID(l.ctx, in.Page, in.Size, in.AuthorID, params.ToOrderByOpt(in.OrderBy))
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] GetPostsByAuthorID error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfos := make([]*post.PostInfo, 0, len(posts))
	for _, po := range posts {
		poInfo := converter.PostDBToRPC(po)
		poInfos = append(poInfos, poInfo)
	}

	return &post.GetPostsByAuthorIDResponse{Posts: poInfos, Paginator: paginator.ToRPC()}, nil
}
