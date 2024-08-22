package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/service/post/rpc/internal/svc"
	"lookingforpartner/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/core/logx"
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
		Logger: logx.WithContext(ctx),
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
	updatedProj, err := l.svcCtx.PostInterface.SetProject(&proj)
	if err != nil {
		l.Logger.Errorf("[Post][Rpc] SetProject error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	projResp := converter.Project2ProjResp(updatedProj)

	return &post.UpdateProjectResponse{Project: projResp}, nil
}
