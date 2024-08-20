package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/api/internal/converter"
	"lookingforpartner/service/post/rpc/pb/post"

	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		Logger: logx.WithContext(ctx),
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

	resp = &types.GetPostResponse{
		Post: converter.PostRpc2Api(getPostResp.GetPost()),
	}

	return resp, nil
}
