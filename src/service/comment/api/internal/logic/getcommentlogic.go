package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/api/internal/converter"
	"lookingforpartner/service/comment/api/internal/svc"
	"lookingforpartner/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentRequest) (resp *types.GetCommentResponse, err error) {
	getCommentReq := comment.GetCommentRequest{CommentID: req.CommentID}
	getCommentResp, err := l.svcCtx.CommentRpc.GetComment(l.ctx, &getCommentReq)
	if err != nil {
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.GetCommentResponse{Comment: converter.CommentRpcToApi(getCommentResp.Comment)}

	return resp, nil
}
