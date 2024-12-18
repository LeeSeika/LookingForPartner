package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/logic"
	"lookingforpartner/service/comment/rpc/internal/svc"
)

type DeletePost struct {
	svcCtx *svc.ServiceContext
	conf   config.Config
	logger logx.Logger
}

func NewDeletePost(conf config.Config, svcCtx *svc.ServiceContext) *DeletePost {
	return &DeletePost{
		svcCtx: svcCtx,
		conf:   conf,
		logger: logx.WithContext(context.Background()).WithFields(logx.Field("consumer-topic", conf.KqDeletePostConsumerConf.Topic)),
	}
}

func (c *DeletePost) Consume(ctx context.Context, key, val string) error {
	var evt event.DeletePost
	if err := json.Unmarshal([]byte(val), &evt); err != nil {
		c.logger.
			Errorf("cannot unmarshal json, err: %+v", err)
		return err
	}

	subject, err := c.svcCtx.CommentInterface.GetSubjectByPostID(ctx, evt.Post.PostID)
	if err != nil {
		c.logger.Errorf("cannot get subject by post id, err: %+v", err)
		return err
	}

	deleteSubjectReq := comment.DeleteSubjectRequest{SubjectID: subject.SubjectID}
	deleteSubjectLogic := logic.NewDeleteSubjectLogic(ctx, c.svcCtx)
	deletedSubject, err := deleteSubjectLogic.DeleteSubject(&deleteSubjectReq)
	if err != nil {
		c.logger.Errorf("cannot delete subject, err: %+v", err)
		return err
	}

	deleteAllCommentsLogic := logic.NewDeleteAllCommentsBySubjectIDLogic(ctx, c.svcCtx)
	// no need to handle error
	_, _ = deleteAllCommentsLogic.DeleteAllCommentsBySubjectID(&comment.DeleteAllCommentsBySubjectIDRequest{SubjectID: deletedSubject.Subject.SubjectID})

	return nil
}
