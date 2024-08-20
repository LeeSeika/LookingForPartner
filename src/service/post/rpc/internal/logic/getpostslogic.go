package logic

import (
	"context"
	"lookingforpartner/common/dao"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/service/post/rpc/internal/svc"
	"lookingforpartner/service/post/rpc/pb/post"

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
	posts, err := l.svcCtx.PostInterface.GetPosts(in.Page, in.Size, dao.OrderByString2Opt(in.OrderBy))
	if err != nil {
		return nil, errs.RpcUnknown
	}
	if len(posts) == 0 {
		return nil, errs.RpcNotFound
	}

	poInfos := make([]*post.PostInfo, 0, len(posts))
	for _, po := range posts {
		poInfo := converter.PostWithProject2PostInfo(po)
		poInfos = append(poInfos, poInfo)
	}

	return &post.GetPostsResponse{Posts: poInfos}, nil
}