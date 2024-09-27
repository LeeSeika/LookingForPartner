package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/model/entity"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FillSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFillSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FillSubjectLogic {
	return &FillSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FillSubjectLogic) FillSubject(in *post.FillSubjectRequest) (*post.FillSubjectResponse, error) {

	updatedPost := entity.Post{
		PostID:           in.PostID,
		CommentSubjectID: in.SubjectID,
	}
	_, err := l.svcCtx.PostInterface.UpdatePost(l.ctx, &updatedPost)
	if err != nil {
		l.Logger.Errorf("cannot update post, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &post.FillSubjectResponse{}, nil
}
