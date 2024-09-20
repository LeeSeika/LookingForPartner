package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/comment"
	"lookingforpartner/pb/paginator"
	"lookingforpartner/service/comment/api/internal/converter"

	"lookingforpartner/service/comment/api/internal/svc"
	"lookingforpartner/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByPostIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPostIDLogic {
	return &GetCommentsByPostIDLogic{
		Logger: logger.NewLogger(ctx, "comment-api"),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsByPostIDLogic) GetCommentsByPostID(req *types.GetCommentsByPostIDRequest) (resp *types.GetCommentsByPostIDResponse, err error) {
	getCommentsByPostIDReq := comment.GetCommentsByPostIDRequest{
		PostID: req.PostID,
		PaginationParams: &paginator.PaginationParams{
			Page:    req.PaginationParams.Page,
			Size:    req.PaginationParams.Size,
			OrderBy: req.PaginationParams.Order,
		},
	}
	getCommentsByPostIDResp, err := l.svcCtx.CommentRpc.GetCommentsByPostID(l.ctx, &getCommentsByPostIDReq)
	if err != nil {
		return nil, errs.FormattedApiInternal()
	}

	commentRpcs := getCommentsByPostIDResp.Comments
	pagi := getCommentsByPostIDResp.Paginator

	commentApis := make([]types.Comment, 0, len(commentRpcs))

	for i := 0; i < len(commentRpcs); i++ {
		commentApi := converter.CommentRpcToApi(commentRpcs[i])
		commentApis = append(commentApis, commentApi)
	}

	resp = &types.GetCommentsByPostIDResponse{Comments: commentApis, Paginator: types.Paginator{
		TotalRecord: pagi.TotalRecord,
		TotalPage:   int(pagi.TotalPage),
		Offset:      int(pagi.Offset),
		Limit:       int(pagi.Limit),
		CurrPage:    int(pagi.CurrPage),
		PrevPage:    int(pagi.PrevPage),
		NextPage:    int(pagi.NextPage),
	}}

	return resp, nil
}
