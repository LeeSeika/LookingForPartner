package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/api/internal/converter"

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

	createCommentReq := comment.CreateCommentRequest{
		WechatID:  uid,
		SubjectID: req.SubjectID,
		RootID:    req.RootID,
		ParentID:  req.ParentID,
		DialogID:  req.DialogID,
		Content:   req.Content,
	}

	createCommentResp, err := l.svcCtx.CommentRpc.CreateComment(l.ctx, &createCommentReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.CreateCommentResponse{Comment: converter.CommentRpcToApi(createCommentResp.Comment)}

	return resp, nil
}
