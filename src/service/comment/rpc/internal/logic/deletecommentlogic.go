package logic

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/comment"
	"lookingforpartner/pb/post"
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
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {

	_comment, err := l.svcCtx.CommentInterface.GetComment(l.ctx, in.CommentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get comment when deleting comment, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}
	subject, err := l.svcCtx.CommentInterface.GetSubject(l.ctx, _comment.SubjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get subject when deleting comment, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	// check permission
	if _comment.AuthorID != in.OperatorID {
		// get post and check post author id
		getPostReq := post.GetPostRequest{PostID: subject.PostID}
		getPostResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, &getPostReq)
		if err != nil {
			l.Logger.Errorf("cannot get post when deleting comment, err: %+v", err)
			return nil, err
		}

		postAuthorID := getPostResp.Post.Author.WxUid

		if postAuthorID != in.OperatorID {
			return nil, errs.RpcPermissionDenied
		}
	}

	// delete comment
	deletedComment, err := l.svcCtx.CommentInterface.DeleteComment(l.ctx, in.CommentID)
	if err != nil {
		l.Logger.Errorf("cannot delete comment, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	// if this is a root comment, asynchronously delete all of its sub comments
	if deletedComment.RootID == nil {
		evt := event.DeleteRootComment{CommentID: in.CommentID}
		bytes, err := json.Marshal(&evt)
		if err != nil {
			l.Logger.Errorf("cannot marshal event when deleting comment, err: %+v", err)
			return nil, errs.FormatRpcUnknownError(err.Error())
		}

		err = l.svcCtx.KqDeleteRootComment.Push(l.ctx, string(bytes))
		if err != nil {
			topic := l.svcCtx.Config.KqDeleteRootCommentPusherConf.Topic
			l.Logger.
				WithFields(logx.Field("topic", topic)).
				Errorf("cannot push a message to mq when deleting comment, err: %+V", err)

			l.svcCtx.LocalQueue.Push(evt)
		}
	}

	return &comment.DeleteCommentResponse{}, nil
}
