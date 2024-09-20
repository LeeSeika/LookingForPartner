package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/comment"
	"net/http"

	"lookingforpartner/service/comment/api/internal/svc"
	"lookingforpartner/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logger.NewLogger(ctx, "comment-api"),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentRequest) (resp *types.DeleteCommentResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}

	deleteCommentReq := comment.DeleteCommentRequest{
		CommentID:  req.CommentID,
		OperatorID: uid,
	}

	_, err = l.svcCtx.CommentRpc.DeleteComment(l.ctx, &deleteCommentReq)
	if err != nil {
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		} else if errors.Is(err, errs.RpcPermissionDenied) {
			return nil, errs.FormatApiError(http.StatusForbidden, errs.ApiPermissionDenied)
		}

		return nil, errs.FormattedApiInternal()
	}

	resp = &types.DeleteCommentResponse{}

	return resp, nil
}
