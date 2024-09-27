package mqs

import (
	"context"
	"lookingforpartner/common/constant"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/logic"
	"lookingforpartner/service/comment/rpc/internal/svc"
)

type DeleteCommentsByID struct {
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentsByID(svcCtx *svc.ServiceContext) *DeleteCommentsByID {
	return &DeleteCommentsByID{
		svcCtx: svcCtx,
	}
}

func (c *DeleteCommentsByID) Consume(ctx context.Context, key, val string) error {
	switch key {
	case constant.MqMessageKeyDeleteAllCommentsBySubjectID:
		return c.deleteAllCommentsBySubjectID(ctx, val)

	case constant.MqMessageKeyDeleteSubCommentsByRootID:
		return c.deleteSubCommentsByRootID(ctx, val)

	default:
		return nil
	}
}

func (c *DeleteCommentsByID) deleteSubCommentsByRootID(ctx context.Context, rootID string) error {
	DeleteSubCommentsByRootIDReq := comment.DeleteSubCommentsByRootIDRequest{RootID: rootID}

	l := logic.NewDeleteSubCommentsByRooIDLogic(ctx, c.svcCtx)
	_, err := l.DeleteSubCommentsByRooID(&DeleteSubCommentsByRootIDReq)
	if err != nil {
		l.Logger.Errorf("cannot delete sub comments by root id, err: %+v", err)
		return err
	}

	return nil
}

func (c *DeleteCommentsByID) deleteAllCommentsBySubjectID(ctx context.Context, subjectID string) error {
	DeleteAllCommentsBySubjectIDReq := comment.DeleteAllCommentsBySubjectIDRequest{SubjectID: subjectID}

	l := logic.NewDeleteAllCommentsBySubjectIDLogic(ctx, c.svcCtx)
	_, err := l.DeleteAllCommentsBySubjectID(&DeleteAllCommentsBySubjectIDReq)
	if err != nil {
		l.Logger.Errorf("cannot delete all comments by subject id, err: %+v", err)
		return err
	}

	return nil
}
