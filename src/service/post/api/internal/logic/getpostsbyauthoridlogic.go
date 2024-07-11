package logic

import (
	"context"

	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsByAuthorIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsByAuthorIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsByAuthorIDLogic {
	return &GetPostsByAuthorIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsByAuthorIDLogic) GetPostsByAuthorID(req *types.GetPostByAuthorIDRequest) (resp *types.GetPostsByAuthorIDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
