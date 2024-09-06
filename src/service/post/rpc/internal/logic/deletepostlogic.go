package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/user"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
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
		Logger: logger.NewLogger(ctx, "post"),
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
	if po.AuthorID != in.WxUid {
		return nil, errs.RpcPermissionDenied
	}
	_, err = l.svcCtx.PostInterface.DeletePost(l.ctx, in.PostID, in.IdempotencyKey)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			//
			return nil, errs.RpcAlreadyExists
		}
		l.Logger.Errorf("cannot delete post, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	// call user rpc
	updateUserPostCountReq := user.UpdateUserPostCountRequest{
		IdempotencyKey: in.IdempotencyKey,
		WxUid:          in.WxUid,
		Delta:          -1,
	}
	_, err = l.svcCtx.UserRpc.UpdateUserPostCount(l.ctx, &updateUserPostCountReq)
	if err != nil {
		// push to kafka to retry asynchronously
		bytes, err := json.Marshal(&updateUserPostCountReq)
		if err != nil {
			// todo: add local queue
		}

		err = l.svcCtx.KqUpdateUserPostCountPusher.Push(l.ctx, string(bytes))
		if err != nil {
			// todo: add local queue
		}
	}

	return &post.DeletePostResponse{}, nil
}
