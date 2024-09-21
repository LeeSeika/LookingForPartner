package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/post"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/rpc/internal/converter"
	"lookingforpartner/service/post/rpc/internal/svc"
)

type GetPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "post-rpc"),
	}
}

func (l *GetPostLogic) GetPost(in *post.GetPostRequest) (*post.GetPostResponse, error) {
	poProj, err := l.svcCtx.PostInterface.GetPost(l.ctx, in.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get post, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	poInfo := converter.PostDBToRPC(poProj.Post)
	if poProj.Project != nil {
		projInfo := converter.ProjectDBToRPC(poProj.Project)
		poInfo.Project = projInfo
	}

	// get author & maintainer info from user
	getUserInfoReq := user.GetUserInfoRequest{WxUid: poProj.AuthorID}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		l.Logger.Errorf("cannot get author info when getting post, err:%+v", err)
	} else {
		userInfo := getUserInfoResp.UserInfo
		poInfo.Author = userInfo
		if poInfo.Project != nil {
			poInfo.Project.Maintainer = userInfo
		}
	}

	return &post.GetPostResponse{Post: poInfo}, nil
}
