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

type UpdateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProjectLogic) UpdateProject(req *types.UpdateProjectRequest) (resp *types.UpdateProjectResponse, err error) {
	updateProjectReq := post.UpdateProjectRequest{
		ProjectID:     req.ProjectID,
		Name:          req.Name,
		Role:          req.Role,
		Introduction:  req.Introduction,
		Progress:      req.Progress,
		HeadCountInfo: req.HeadCountInfo,
	}

	updateProjectResp, err := l.svcCtx.PostRpc.UpdateProject(l.ctx, &updateProjectReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.UpdateProjectResponse{
		Project: converter.ProjectRpcToApi(updateProjectResp.GetProject()),
	}

	return resp, nil
}
