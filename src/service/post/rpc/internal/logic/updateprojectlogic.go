package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/post/model/entity"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type UpdateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProjectLogic {
	return &UpdateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "post-rpc"),
	}
}

func (l *UpdateProjectLogic) UpdateProject(in *post.UpdateProjectRequest) (*post.UpdateProjectResponse, error) {
	proj := entity.Project{
		ProjectID:     in.ProjectID,
		Name:          in.Name,
		Introduction:  in.Introduction,
		Role:          in.Role,
		HeadCountInfo: in.HeadCountInfo,
		Progress:      in.Progress,
	}
	updatedProj, err := l.svcCtx.PostInterface.UpdateProject(l.ctx, &proj)
	if err != nil {
		l.Logger.Errorf("cannot update project, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	projResp := converter.ProjectDBToRPC(updatedProj)

	return &post.UpdateProjectResponse{Project: projResp}, nil
}
