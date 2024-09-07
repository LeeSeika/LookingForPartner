package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/comment"

	"lookingforpartner/service/comment/api/internal/svc"
	"lookingforpartner/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentRequest) (resp *types.CreateCommentResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}

	createCommentReq := comment.CreateCommentRequest{}

	return
}
