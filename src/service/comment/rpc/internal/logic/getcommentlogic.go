package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/comment/rpc/internal/converter"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

func (l *GetCommentLogic) GetComment(in *comment.GetCommentRequest) (*comment.GetCommentResponse, error) {

	_commentIndexContent, err := l.svcCtx.CommentInterface.GetComment(l.ctx, in.CommentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get comment, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &comment.GetCommentResponse{Comment: converter.SingleCommentDBToRPC(_commentIndexContent.CommentIndex, _commentIndexContent.CommentContent)}, nil
}
