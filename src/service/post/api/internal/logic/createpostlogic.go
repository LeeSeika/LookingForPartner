package logic

import (
	"context"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/api/internal/converter"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostRequest) (resp *types.CreatePostResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}

	proj := converter.ProjectApiToRpc(&req.Project)
	createPostReq := post.CreatePostRequest{
		Title:   req.Title,
		Project: &proj,
		Content: req.Content,
		WxUid:   uid,
	}
	createPostResp, err := l.svcCtx.PostRpc.CreatePost(l.ctx, &createPostReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.CreatePostResponse{
		Post: converter.PostRpcToApi(createPostResp.GetPostInfo()),
	}

	return resp, nil
}
