package logic

import (
	"context"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPostIDLogic {
	return &GetCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsByPostIDLogic) GetCommentsByPostID(in *comment.GetCommentsByPostIDRequest) (*comment.GetCommentsByPostIDResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.GetCommentsByPostIDResponse{}, nil
}
