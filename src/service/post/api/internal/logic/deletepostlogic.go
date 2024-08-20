package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/rpc/pb/post"
	"net/http"

	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
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
			return nil, errs.FormatApiError(http.StatusForbidden, "no permission to delete")
		}
		return nil, errs.FormattedApiInternal()
	}

	return &types.DeletePostResponse{}, nil
}
