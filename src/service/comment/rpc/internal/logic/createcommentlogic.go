package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/pkg/nanoid"
	"lookingforpartner/service/comment/model/entity"
	"lookingforpartner/service/comment/rpc/internal/converter"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	_, err := l.svcCtx.CommentInterface.GetSubject(l.ctx, in.SubjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		l.Logger.Errorf("cannot get subject, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}
	commentID := constant.NanoidPrefixComment + nanoid.Gen()
	commentIndex := &entity.CommentIndex{
		CommentID:       commentID,
		SubjectID:       in.SubjectID,
		RootID:          &in.RootID,
		ParentID:        &in.ParentID,
		DialogID:        &in.DialogID,
		AuthorID:        in.WechatID,
		SubCount:        0,
		LikeCount:       0,
		SubCommentCount: 0,
		Status:          0,
	}

	commentContent := &entity.CommentContent{
		CommentID: commentID,
		Content:   in.Content,
		MetaData:  nil,
	}

	commentIndexContent, err := l.svcCtx.CommentInterface.CreateComment(l.ctx, commentIndex, commentContent)
	if err != nil {
		l.Logger.Errorf("cannot create comment, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &comment.CreateCommentResponse{Comment: converter.SingleCommentDBToRPC(commentIndexContent.CommentIndex, commentIndexContent.CommentContent)}, nil
}
