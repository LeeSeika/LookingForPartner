package logic

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/api/internal/converter"
	"lookingforpartner/service/post/rpc/pb/post"
	"net/http"

	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func validateGetPostsRequest(req *types.GetPostsRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Page, validation.Min(1)),
		validation.Field(&req.Page, validation.Min(-1)),
	)
}

func (l *GetPostsLogic) GetPosts(req *types.GetPostsRequest) (resp *types.GetPostsResponse, err error) {
	if err := validateGetPostsRequest(req); err != nil {
		return nil, errs.FormatApiError(http.StatusBadRequest, err.Error())
	}

	getPostsReq := post.GetPostsRequest{
		Page:    req.Page,
		Size:    req.Size,
		OrderBy: req.Order,
	}

	getPostsResp, err := l.svcCtx.PostRpc.GetPosts(l.ctx, &getPostsReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	posts := getPostsResp.GetPosts()
	postInfos := make([]types.Post, 0, len(posts))
	for _, poRpc := range posts {
		poApi := converter.PostRpc2Api(poRpc)
		postInfos = append(postInfos, poApi)
	}

	resp = &types.GetPostsResponse{Posts: postInfos}

	return resp, nil
}
