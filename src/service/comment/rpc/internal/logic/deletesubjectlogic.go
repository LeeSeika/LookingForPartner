package logic

import (
	"context"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/comment/model/dto"

	"lookingforpartner/pb/comment"
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
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

func (l *DeleteSubjectLogic) DeleteSubject(in *comment.DeleteSubjectRequest) (*comment.DeleteSubjectResponse, error) {
	deletedSubject, err := l.svcCtx.CommentInterface.DeleteSubject(l.ctx, in.SubjectID)
	if err != nil {
		l.Logger.Errorf("cannot delete subject, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	// asynchronously delete all comments of this subject
	err = l.svcCtx.KqDeleteCommentsByIDPusher.KPush(l.ctx, constant.MqMessageKeyDeleteAllCommentsBySubjectID, deletedSubject.SubjectID)
	if err != nil {
		topic := l.svcCtx.Config.KqDeleteCommentsByIDPusherConf.Topic
		l.Logger.
			WithFields(logx.Field("topic", topic)).
			WithFields(logx.Field("key", constant.MqMessageKeyDeleteAllCommentsBySubjectID)).
			Errorf("cannot push a message to mq when deleting subject, err: %+V", err)

		msg := dto.DeleteCommentMessage{
			Topic: topic,
			Key:   constant.MqMessageKeyDeleteAllCommentsBySubjectID,
			Val:   deletedSubject.SubjectID,
		}
		l.svcCtx.LocalQueue.Push(msg)
	}

	return &comment.DeleteSubjectResponse{}, nil
}
