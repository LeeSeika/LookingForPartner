package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/common/params"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type GetPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "post-rpc"),
	}
}

func (l *GetPostsLogic) GetPosts(in *post.GetPostsRequest) (*post.GetPostsResponse, error) {
	poProjs, paginator, err := l.svcCtx.PostInterface.GetPosts(l.ctx, in.PaginationParams.Page, in.PaginationParams.Size, params.ToOrderByOpt(in.PaginationParams.OrderBy))
	if err != nil {
		l.Logger.Errorf("cannot get posts, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	poInfos := make([]*post.PostInfo, 0, len(poProjs))
	for _, poProj := range poProjs {
		poInfo := converter.PostDBToRPC(poProj.Post)
		if poProj.Project != nil {
			projInfo := converter.ProjectDBToRPC(poProj.Project)
			poInfo.Project = projInfo
		}
		poInfos = append(poInfos, poInfo)
	}

	authorIDs := make([]string, 0, len(poInfos))
	authorIDToPoInfoMap := make(map[string]*post.PostInfo, len(authorIDs))
	for i := 0; i < len(poInfos); i++ {
		poInfo := poInfos[i]

		authorIDs = append(authorIDs, poInfo.Author.WxUid)
		authorIDToPoInfoMap[poInfo.Author.WxUid] = poInfo
	}

	// get author & maintainer info from user rpc
	getUserInfoByIDsReq := user.GetUserInfoByIDsRequest{WechatIDs: authorIDs}
	getUserInfoByIDsResp, err := l.svcCtx.UserRpc.GetUserInfoByIDs(l.ctx, &getUserInfoByIDsReq)
	if err != nil {
		l.Logger.Errorf("cannot get author infos when getting posts, err:%+v", err)
	} else {
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

	return &post.GetPostsResponse{Posts: poInfos, Paginator: paginator.ToRPC()}, nil
}
