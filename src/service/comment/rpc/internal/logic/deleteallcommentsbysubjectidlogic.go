package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAllCommentsBySubjectIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAllCommentsBySubjectIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAllCommentsBySubjectIDLogic {
	return &DeleteAllCommentsBySubjectIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

func (l *DeleteAllCommentsBySubjectIDLogic) DeleteAllCommentsBySubjectID(in *comment.DeleteAllCommentsBySubjectIDRequest) (*comment.DeleteAllCommentsBySubjectIDResponse, error) {
	err := l.svcCtx.CommentInterface.DeleteAllCommentsBySubjectID(l.ctx, in.SubjectID)
	if err != nil {
		l.Logger.Errorf("cannot delete all comments by subject id, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &comment.DeleteAllCommentsBySubjectIDResponse{}, nil
}
