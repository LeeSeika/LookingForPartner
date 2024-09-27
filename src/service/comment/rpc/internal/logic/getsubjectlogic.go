package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/comment/rpc/internal/converter"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectLogic {
	return &GetSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

func (l *GetSubjectLogic) GetSubject(in *comment.GetSubjectRequest) (*comment.GetSubjectResponse, error) {

	subject, err := l.svcCtx.CommentInterface.GetSubject(l.ctx, in.SubjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get subject, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &comment.GetSubjectResponse{Subject: converter.SubjectDBToRPC(subject)}, nil
}
