package logic

import (
	"context"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubjectLogic {
	return &CreateSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSubjectLogic) CreateSubject(in *comment.CreateSubjectRequest) (*comment.CreateSubjectResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.CreateSubjectResponse{}, nil
}
