package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPostCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPostCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPostCountLogic {
	return &UpdateUserPostCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPostCountLogic) UpdateUserPostCount(in *user.UpdateUserPostCountRequest) (*user.UpdateUserPostCountResponse, error) {
	err := l.svcCtx.UserInterface.UpdatePostCount(l.ctx, in.WxUid, int(in.Delta), in.IdempotencyKey)
	if err != nil {
		l.Logger.Errorf("cannot update user post count, err: %+v", err)
		if errors.Is(err, errs.DBDuplicatedIdempotencyKey) {
			return nil, errs.RpcDuplicatedIdempotencyKey
		}
		return nil, errs.RpcUnknown
	}

	return &user.UpdateUserPostCountResponse{}, nil
}
