package logic

import (
	"context"
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

type GetPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{

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
		PaginationParams: &paginator.PaginationParams{
			Page:    req.Page,
			Size:    req.Size,
			OrderBy: req.Order,
		},
	}

	getPostsResp, err := l.svcCtx.PostRpc.GetPosts(l.ctx, &getPostsReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	posts := getPostsResp.GetPosts()

	authorIDs := make([]string, 0, len(posts))
	authorIDToPoInfoMap := make(map[string]*post.PostInfo, len(authorIDs))
	for i := 0; i < len(posts); i++ {
		poInfo := posts[i]

		authorIDs = append(authorIDs, poInfo.Author.WxUid)
		authorIDToPoInfoMap[poInfo.Author.WxUid] = poInfo
	}

	// get authors info
	getUserInfoByIDsReq := user.GetUserInfoByIDsRequest{WechatIDs: authorIDs}
	getUserInfoByIDsResp, err := l.svcCtx.UserRpc.GetUserInfoByIDs(l.ctx, &getUserInfoByIDsReq)
	if err == nil {
		userInfos := getUserInfoByIDsResp.UserInfos
		for i := 0; i < len(userInfos); i++ {
			userInfo := userInfos[i]
			authorID := userInfo.WxUid

			poInfo := authorIDToPoInfoMap[authorID]
			poInfo.Author = userInfo
			if poInfo.Project != nil {
				poInfo.Project.Maintainer = userInfo
			}
		}
	}

	postInfos := make([]types.Post, 0, len(posts))
	for _, poRpc := range posts {
		poApi := converter.PostRpcToApi(poRpc)
		postInfos = append(postInfos, poApi)
	}

	resp = &types.GetPostsResponse{Posts: postInfos, Paginator: types.Paginator{
		TotalRecord: getPostsResp.Paginator.TotalRecord,
		TotalPage:   int(getPostsResp.Paginator.TotalPage),
		Offset:      int(getPostsResp.Paginator.Offset),
		Limit:       int(getPostsResp.Paginator.Limit),
		CurrPage:    int(getPostsResp.Paginator.CurrPage),
		PrevPage:    int(getPostsResp.Paginator.PrevPage),
		NextPage:    int(getPostsResp.Paginator.NextPage),
	}}

	return resp, nil
}
