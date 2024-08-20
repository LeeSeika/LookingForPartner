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
	posts, err := l.svcCtx.PostInterface.GetPostsByAuthorID(in.Page, in.Size, in.AuthorID, dao.OrderByString2Opt(in.OrderBy))
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] GetPostsByAuthorID error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfos := make([]*post.PostInfo, 0, len(posts))
	for _, po := range posts {
		poInfo := converter.PostWithProject2PostInfo(po)
		poInfos = append(poInfos, poInfo)
	}

	return &post.GetPostsByAuthorIDResponse{Posts: poInfos}, nil
}
