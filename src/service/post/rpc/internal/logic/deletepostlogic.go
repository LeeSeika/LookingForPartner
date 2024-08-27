package logic

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(in *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	po, err := l.svcCtx.PostInterface.GetPost(l.ctx, in.PostID)
	if err != nil {
		if po == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		log.Error().Msgf("cannot get post, err: %+v", err)
		return nil, errs.RpcUnknown
	}
	if po.AuthorID != in.WxUid {
		return nil, errs.RpcPermissionDenied
	}
	_, err = l.svcCtx.PostInterface.DeletePost(l.ctx, in.PostID)
	if err != nil {
		log.Error().Msgf("cannot delete post, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &post.DeletePostResponse{}, nil
}
