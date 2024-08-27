package logic

import (
	"context"
	"github.com/rs/zerolog/log"
	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
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

func (l *UpdateProjectLogic) UpdateProject(in *post.UpdateProjectRequest) (*post.UpdateProjectResponse, error) {
	proj := model.Project{
		ProjectID:     in.ProjectID,
		Name:          in.Name,
		Introduction:  in.Introduction,
		Role:          in.Role,
		HeadCountInfo: in.HeadCountInfo,
		Progress:      in.Progress,
	}
	updatedProj, err := l.svcCtx.PostInterface.UpdateProject(l.ctx, &proj)
	if err != nil {
		log.Error().Msgf("cannot update project, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	projResp := converter.ProjectDBToRPC(updatedProj)

	return &post.UpdateProjectResponse{Project: projResp}, nil
}
