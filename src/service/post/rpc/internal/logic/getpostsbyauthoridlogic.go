package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/common/params"
	"lookingforpartner/service/post/rpc/internal/converter"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type GetPostsByAuthorIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostsByAuthorIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsByAuthorIDLogic {
	return &GetPostsByAuthorIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "post"),
	}
}

func (l *GetPostsByAuthorIDLogic) GetPostsByAuthorID(in *post.GetPostsByAuthorIDRequest) (*post.GetPostsByAuthorIDResponse, error) {
	poProjs, paginator, err := l.svcCtx.PostInterface.GetPostsByAuthorID(l.ctx, in.Page, in.Size, in.AuthorID, params.ToOrderByOpt(in.OrderBy))
	if err != nil {
		l.Logger.Errorf("cannot get posts by author_id, err: %+v", err)
		return nil, errs.RpcUnknown
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

	return &post.GetPostsByAuthorIDResponse{Posts: poInfos, Paginator: paginator.ToRPC()}, nil
}
