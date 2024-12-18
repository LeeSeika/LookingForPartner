package logic

import (
	"context"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/post"
	"lookingforpartner/pkg/nanoid"
	"lookingforpartner/service/comment/model/entity"
	"lookingforpartner/service/comment/rpc/internal/converter"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubjectLogic {
	return &CreateSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSubjectLogic) CreateSubject(in *comment.CreateSubjectRequest) (*comment.CreateSubjectResponse, error) {
	subjectID := constant.NanoidPrefixSubject + nanoid.Gen()
	subject := &entity.Subject{
		SubjectID:        subjectID,
		PostID:           in.PostID,
		AllCommentCount:  0,
		RootCommentCount: 0,
	}

	subject, err := l.svcCtx.CommentInterface.CreateSubject(l.ctx, subject, in.IdempotencyKey)
	if err != nil {
		l.Logger.Errorf("cannot create subject, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	// fill comment subject id
	fillSubjectReq := post.FillSubjectRequest{
		PostID:    in.PostID,
		SubjectID: subjectID,
	}
	_, err = l.svcCtx.PostRpc.FillSubject(l.ctx, &fillSubjectReq)
	if err != nil {
		l.Logger.Errorf("cannot fill subject id, err: %+v", err)
		// todo: retry
	}

	return &comment.CreateSubjectResponse{Subject: converter.SubjectDBToRPC(subject)}, nil
}
