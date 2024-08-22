package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	model2 "lookingforpartner/model"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"
	"lookingforpartner/service/post/rpc/pb/post"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	var po *model2.Post
	var proj *model2.Project
	var err error

	po = &model2.Post{
		PostID:   1, //todo
		AuthorID: in.WxUid,
		Title:    in.Title,
		Content:  in.Content,
	}

	if in.Project == nil {
		po, err = l.svcCtx.PostInterface.CreatePost(po)
	} else {
		proj = &model2.Project{
			ProjectID:     1, //todo
			MaintainerID:  in.Project.MaintainerID,
			Maintainer:    in.Project.Maintainer,
			Name:          in.Project.Name,
			Introduction:  in.Project.Introduction,
			Role:          in.Project.Role,
			HeadCountInfo: in.Project.HeadCountInfo,
			Progress:      in.Project.Progress,
			PostID:        po.PostID,
		}
		po, proj, err = l.svcCtx.PostInterface.CreatePostWithProjectTx(po, proj)
	}

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, errs.RpcAlreadyExists
		}
		l.Logger.Errorf("[Post][Rpc] CreatePostWithProjectTx error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfo := converter.PostAndProject2PostInfo(po, proj)

	return &post.CreatePostResponse{PostInfo: poInfo}, nil
}
