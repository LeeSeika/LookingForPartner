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

type GetPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostsLogic) GetPosts(in *post.GetPostsRequest) (*post.GetPostsResponse, error) {
	posts, paginator, err := l.svcCtx.PostInterface.GetPosts(l.ctx, in.Page, in.Size, params.ToOrderByOpt(in.OrderBy))
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] GetPosts error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfos := make([]*post.PostInfo, 0, len(posts))
	for _, po := range posts {
		poInfo := converter.PostDBToRPC(po)
		poInfos = append(poInfos, poInfo)
	}

	return &post.GetPostsResponse{Posts: poInfos, Paginator: paginator.ToRPC()}, nil
}
