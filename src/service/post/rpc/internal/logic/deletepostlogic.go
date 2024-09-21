package logic

import (
	"context"
	"encoding/json"
	"errors"
	"lookingforpartner/pb/comment"

	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/post"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/model/dto"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "post-rpc"),
	}
}

func (l *DeletePostLogic) DeletePost(in *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	po, err := l.svcCtx.PostInterface.GetPost(l.ctx, in.PostID)
	if err != nil {
		if po == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get post, err: %+v", err)
		return nil, errs.RpcUnknown
	}
	// check permission
	if po.AuthorID != in.WxUid {
		return nil, errs.RpcPermissionDenied
	}
	_, err = l.svcCtx.PostInterface.DeletePost(l.ctx, in.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			//
			return nil, errs.RpcAlreadyExists
		}
		l.Logger.Errorf("cannot delete post, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	// update user post count
	updateUserPostCountReq := user.UpdateUserPostCountRequest{
		IdempotencyKey: 123, // todo: gen idem key
		WxUid:          in.WxUid,
		Delta:          -1,
	}
	_, err = l.svcCtx.UserRpc.UpdateUserPostCount(l.ctx, &updateUserPostCountReq)
	if err != nil {
		// push to kafka to retry asynchronously
		bytes, _ := json.Marshal(&updateUserPostCountReq)

		err = l.svcCtx.KqUpdateUserPostCountPusher.Push(l.ctx, string(bytes))
		if err != nil {
			// push to local queue
			topic := l.svcCtx.Config.KqUpdateUserPostCountPusherConf.Topic
			l.Logger.
				WithFields(logx.Field("topic", topic)).
				Errorf("cannot push a message to mq when deleting post, err: %+V", err)

			msg := dto.UpdateUserPostCountMessage{
				Topic: topic,
				Val:   bytes,
			}
			l.svcCtx.LocalQueue.Push(msg)
		}
	}

	// delete comment subject
	deleteSubjectReq := comment.DeleteSubjectRequest{
		SubjectID: po.CommentSubjectID,
	}
	_, err = l.svcCtx.CommentRpc.DeleteSubject(l.ctx, &deleteSubjectReq)
	if err != nil {
		l.Logger.Errorf("cannot delete subject when deleting post, err: %+v", err)
		// push to local queue to retry
		msg := dto.DeleteSubjectMessage{
			Topic:     l.svcCtx.Config.KqDeleteSubjectPusherConf.Topic,
			SubjectID: po.CommentSubjectID,
		}
		l.svcCtx.LocalQueue.Push(msg)
	}

	return &post.DeletePostResponse{}, nil
}
