package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"

	"lookingforpartner/service/post/rpc/internal/svc"
	"lookingforpartner/service/post/rpc/pb/post"

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
	po, err := l.svcCtx.PostInterface.GetPost(in.PostID)
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
	_, _, err = l.svcCtx.PostInterface.DeletePostTx(in.PostID)
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] CreatePostWithProjectTx error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &post.DeletePostResponse{}, nil
}
