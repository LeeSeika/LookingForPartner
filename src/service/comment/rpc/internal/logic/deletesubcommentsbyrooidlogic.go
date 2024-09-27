package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubCommentsByRooIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSubCommentsByRooIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubCommentsByRooIDLogic {
	return &DeleteSubCommentsByRooIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

func (l *DeleteSubCommentsByRooIDLogic) DeleteSubCommentsByRooID(in *comment.DeleteSubCommentsByRootIDRequest) (*comment.DeleteSubjectResponse, error) {
	err := l.svcCtx.CommentInterface.DeleteSubCommentsByRootID(l.ctx, in.RootID)
	if err != nil {
		l.Logger.Errorf("cannot delete sub comments by root id, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &comment.DeleteSubjectResponse{}, nil
}
