package logic

import (
	"context"

	"lookingforpartner/service/comment/api/internal/svc"
	"lookingforpartner/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentRequest) (resp *types.CreateCommentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
