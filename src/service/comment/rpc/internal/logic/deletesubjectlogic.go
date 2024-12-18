package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/converter"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectLogic {
	return &DeleteSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSubjectLogic) DeleteSubject(in *comment.DeleteSubjectRequest) (*comment.DeleteSubjectResponse, error) {
	deletedSubject, err := l.svcCtx.CommentInterface.DeleteSubject(l.ctx, in.SubjectID)
	if err != nil {
		l.Logger.Errorf("cannot delete subject, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &comment.DeleteSubjectResponse{Subject: converter.SubjectDBToRPC(deletedSubject)}, nil
}
