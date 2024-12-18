package logic

import (
	"context"
	"errors"
	"lookingforpartner/pb/user"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/api/internal/converter"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"
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

func (l *GetPostLogic) GetPost(req *types.GetPostRequest) (resp *types.GetPostResponse, err error) {
	getPostReq := post.GetPostRequest{
		PostID: req.PostID,
	}

	getPostResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, &getPostReq)
	if err != nil {
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	// get author info
	getUserInfoReq := user.GetUserInfoRequest{
		WxUid: getPostResp.GetPost().GetAuthor().WxUid,
	}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	po := getPostResp.GetPost()
	po.Author = getUserInfoResp.UserInfo

	resp = &types.GetPostResponse{
		Post: converter.PostRpcToApi(po),
	}

	return resp, nil
}
