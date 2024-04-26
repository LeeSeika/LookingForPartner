package handler

import (
	"context"
	"go.uber.org/zap"
	"lookingforpartner/idl/pb/user"
	"lookingforpartner/service/user/rpc/internal/svc"
)

type SignupHandler struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
}

func NewSignupHandler(svcCtx *svc.ServiceContext, ctx context.Context) *SignupHandler {
	return &SignupHandler{
		svcCtx: svcCtx,
		ctx:    ctx,
	}
}

func (l *SignupHandler) Signup(req *user.UserSignupRequest) (*user.UserSignupResponse, error) {
	zap.L().Info("signup")
	return nil, nil
}
