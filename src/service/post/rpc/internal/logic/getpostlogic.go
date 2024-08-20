package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"
	"lookingforpartner/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostLogic) GetPost(in *post.GetPostRequest) (*post.GetPostResponse, error) {
	poWithProj, err := l.svcCtx.PostInterface.GetPost(in.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		return nil, errs.RpcUnknown
	}

	poInfo := converter.PostWithProject2PostInfo(poWithProj)

	return &post.GetPostResponse{Post: poInfo}, nil
}