package logic

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type GetPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostLogic) GetPost(in *post.GetPostRequest) (*post.GetPostResponse, error) {
	po, err := l.svcCtx.PostInterface.GetPost(l.ctx, in.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		log.Error().Msgf("cannot get post, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfo := converter.PostDBToRPC(po)

	return &post.GetPostResponse{Post: poInfo}, nil
}
