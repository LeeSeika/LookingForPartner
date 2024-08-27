package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/api/internal/common"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"

	"github.com/rs/zerolog/log"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{

		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReqeust) (resp *types.RefreshTokenResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}

	// rpc call
	getUserInfoReq := user.GetUserInfoRequest{WxUid: uid}
	_, err = l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		log.Error().Msgf("[User][Api] GetUserInfo error, err: %v", err)
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	// generate token
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshExpire := l.svcCtx.Config.Auth.RefreshExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessToken, refreshToken, err := common.CreateTokenAndRefreshToken(uid, accessExpire, refreshExpire, accessSecret)
	if err != nil {
		log.Error().Msgf("[User][Api] CreateTokenAndRefreshToken error, err: %+v", err)
		return nil, errs.FormattedApiGenTokenFailed()
	}

	return &types.RefreshTokenResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
