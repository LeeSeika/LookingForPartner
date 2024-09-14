package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {

	_comment, err := l.svcCtx.CommentInterface.GetComment(l.ctx, in.CommentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get comment when deleting comment, err: %+v", err)
		return nil, errs.RpcUnknown
	}
	// todo: get post and check post authorid
	// check permission
	if _comment.AuthorID != in.OperatorID {
		return nil, errs.RpcPermissionDenied
	}

	// delete comment
	deletedComment, err := l.svcCtx.CommentInterface.DeleteComment(l.ctx, in.CommentID)
	if err != nil {
		l.Logger.Errorf("cannot delete comment, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	// todo: if this is a root comment, asynchronously delete all of its sub comments
	if deletedComment.RootID == nil {

	}

	return &comment.DeleteCommentResponse{}, nil
}
