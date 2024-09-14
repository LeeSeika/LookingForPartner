package logic

import (
	"context"
	"errors"
	"net/http"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"
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

func (l *DeletePostLogic) DeletePost(req *types.DeletePostRequest) (resp *types.DeletePostResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}

	deletePostReq := post.DeletePostRequest{
		PostID: req.PostID,
		WxUid:  uid,
	}
	_, err = l.svcCtx.PostRpc.DeletePost(l.ctx, &deletePostReq)
	if err != nil {
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		} else if errors.Is(err, errs.RpcPermissionDenied) {
			return nil, errs.FormatApiError(http.StatusForbidden, errs.ApiPermissionDenied)
		}
		return nil, errs.FormattedApiInternal()
	}

	return &types.DeletePostResponse{}, nil
}
