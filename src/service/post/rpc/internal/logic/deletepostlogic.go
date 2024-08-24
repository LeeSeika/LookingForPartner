package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	po, err := l.svcCtx.PostInterface.GetPost(l.ctx, in.PostID)
	if err != nil {
		if po == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("[Post][Rpc] CreatePost error, err: %+v", err)
		return nil, errs.RpcUnknown
	}
	if po.AuthorID != in.WxUid {
		return nil, errs.RpcPermissionDenied
	}
	_, err = l.svcCtx.PostInterface.DeletePost(l.ctx, in.PostID)
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] CreatePostWithProjectTx error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &post.DeletePostResponse{}, nil
}
