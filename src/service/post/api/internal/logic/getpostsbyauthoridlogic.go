package logic

import (
	"context"
	"net/http"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/api/internal/converter"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetPostsByAuthorIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsByAuthorIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsByAuthorIDLogic {
	return &GetPostsByAuthorIDLogic{

		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func validateGetPostByAuthorIDRequest(req *types.GetPostByAuthorIDRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Page, validation.Min(1)),
		validation.Field(&req.Page, validation.Min(-1)),
	)
}

func (l *GetPostsByAuthorIDLogic) GetPostsByAuthorID(req *types.GetPostByAuthorIDRequest) (resp *types.GetPostsByAuthorIDResponse, err error) {
	if err := validateGetPostByAuthorIDRequest(req); err != nil {
		return nil, errs.FormatApiError(http.StatusBadRequest, err.Error())
	}

	getPostsByAuthorIDReq := post.GetPostsByAuthorIDRequest{
		Page:     req.Page,
		Size:     req.Size,
		AuthorID: req.AuthorID,
		OrderBy:  req.Order,
	}

	getPostsByAuthorIDResp, err := l.svcCtx.PostRpc.GetPostsByAuthorID(l.ctx, &getPostsByAuthorIDReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	posts := getPostsByAuthorIDResp.GetPosts()
	postInfos := make([]types.Post, 0, len(posts))
	for _, poRpc := range posts {
		poApi := converter.PostRpcToApi(poRpc)
		postInfos = append(postInfos, poApi)
	}

	resp = &types.GetPostsByAuthorIDResponse{Posts: postInfos}

	return resp, nil
}
