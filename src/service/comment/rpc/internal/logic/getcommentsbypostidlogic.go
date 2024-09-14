package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/common/params"
	"lookingforpartner/service/comment/model/vo"
	"lookingforpartner/service/comment/rpc/internal/converter"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsByPostIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsByPostIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsByPostIDLogic {
	return &GetCommentsByPostIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "comment-rpc"),
	}
}

const DefaultTopCount = 3

func (l *GetCommentsByPostIDLogic) GetCommentsByPostID(in *comment.GetCommentsByPostIDRequest) (*comment.GetCommentsByPostIDResponse, error) {

	rootComments, paginator, err := l.svcCtx.CommentInterface.GetRootCommentsByPostID(l.ctx, in.PostID, in.PaginationParams.Page, in.PaginationParams.Size, params.ToOrderByOpt(in.GetPaginationParams().OrderBy))
	if err != nil {
		l.Logger.Errorf("cannot get root comments, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	rootIDs := make([]string, 0, len(rootComments))
	allSubComments, err := l.svcCtx.CommentInterface.GetTopSubCommentsByRootIDs(l.ctx, rootIDs, DefaultTopCount, params.ToOrderByOpt(params.OrderByCreateTimeASC))
	if err != nil {
		l.Logger.Errorf("cannot get sub comments, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	// construct comment
	subCommentMap := map[string][]*vo.CommentIndexContent{}
	for i := 0; i < len(allSubComments); i++ {
		subComment := allSubComments[i]

		rootID := subComment.CommentIndex.RootID
		subs, ok := subCommentMap[*rootID]
		if !ok {
			subs = make([]*vo.CommentIndexContent, 0, DefaultTopCount)
		}
		subs = append(subs, subComment)

		subCommentMap[*rootID] = subs
	}

	rootCommentInfos := make([]*comment.CommentInfo, 0, len(rootComments))

	for i := 0; i < len(rootComments); i++ {
		rootComment := rootComments[i]

		rootCommentInfo := converter.SingleCommentDBToRPC(rootComment.CommentIndex, rootComment.CommentContent)

		subComments := subCommentMap[rootComment.CommentIndex.CommentID]
		subCommentInfos := make([]*comment.CommentInfo, 0, DefaultTopCount)
		for _, subComment := range subComments {
			subCommentInfo := converter.SingleCommentDBToRPC(subComment.CommentIndex, subComment.CommentContent)
			subCommentInfos = append(subCommentInfos, subCommentInfo)
		}

		rootCommentInfo.SubComments = subCommentInfos
		rootCommentInfos = append(rootCommentInfos, rootCommentInfo)
	}

	return &comment.GetCommentsByPostIDResponse{Comment: rootCommentInfos, Paginator: paginator.ToRPC()}, nil
}
