package logic

import (
	"context"

	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/api/internal/converter"
	"lookingforpartner/service/post/api/internal/svc"
	"lookingforpartner/service/post/api/internal/types"
)

type UpdateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{

		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProjectLogic) UpdateProject(req *types.UpdateProjectRequest) (resp *types.UpdateProjectResponse, err error) {
	updateProjectReq := post.UpdateProjectRequest{
		ProjectID:     req.ProjectID,
		Name:          req.UpdatedProject.Name,
		Role:          req.UpdatedProject.Role,
		Introduction:  req.UpdatedProject.Introduction,
		Progress:      req.UpdatedProject.Progress,
		HeadCountInfo: req.UpdatedProject.HeadCountInfo,
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
