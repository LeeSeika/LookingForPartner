package logic

import (
	"context"
	"errors"
	"lookingforpartner/pb/paginator"
	"lookingforpartner/pb/user"
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
		PaginationParams: &paginator.PaginationParams{
			Page:    req.Page,
			Size:    req.Size,
			OrderBy: req.Order,
		},
		AuthorID: req.AuthorID,
	}

	getPostsByAuthorIDResp, err := l.svcCtx.PostRpc.GetPostsByAuthorID(l.ctx, &getPostsByAuthorIDReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	posts := getPostsByAuthorIDResp.GetPosts()
	if len(posts) == 0 {
		return &types.GetPostsByAuthorIDResponse{}, nil
	}

	// get author info
	getUserInfoReq := user.GetUserInfoRequest{
		WxUid: posts[0].Author.WxUid,
	}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	postInfos := make([]types.Post, 0, len(posts))
	for _, poRpc := range posts {
		poRpc.Author = getUserInfoResp.UserInfo
		poApi := converter.PostRpcToApi(poRpc)
		postInfos = append(postInfos, poApi)
	}

	resp = &types.GetPostsByAuthorIDResponse{Posts: postInfos, Paginator: types.Paginator{
		TotalRecord: getPostsByAuthorIDResp.Paginator.TotalRecord,
		TotalPage:   int(getPostsByAuthorIDResp.Paginator.TotalPage),
		Offset:      int(getPostsByAuthorIDResp.Paginator.Offset),
		Limit:       int(getPostsByAuthorIDResp.Paginator.Limit),
		CurrPage:    int(getPostsByAuthorIDResp.Paginator.CurrPage),
		PrevPage:    int(getPostsByAuthorIDResp.Paginator.PrevPage),
		NextPage:    int(getPostsByAuthorIDResp.Paginator.NextPage),
	}}

	return resp, nil
}
