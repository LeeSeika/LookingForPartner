// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/logic"
	"lookingforpartner/service/comment/rpc/internal/svc"
)

type CommentServer struct {
	svcCtx *svc.ServiceContext
	comment.UnimplementedCommentServer
}

func NewCommentServer(svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) CreateComment(ctx context.Context, in *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

func (s *CommentServer) GetComment(ctx context.Context, in *comment.GetCommentRequest) (*comment.GetCommentResponse, error) {
	l := logic.NewGetCommentLogic(ctx, s.svcCtx)
	return l.GetComment(in)
}

func (s *CommentServer) GetCommentsByPostID(ctx context.Context, in *comment.GetCommentsByPostIDRequest) (*comment.GetCommentsByPostIDResponse, error) {
	l := logic.NewGetCommentsByPostIDLogic(ctx, s.svcCtx)
	return l.GetCommentsByPostID(in)
}

func (s *CommentServer) DeleteComment(ctx context.Context, in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

func (s *CommentServer) DeleteSubCommentsByRooID(ctx context.Context, in *comment.DeleteSubCommentsByRootIDRequest) (*comment.DeleteSubjectResponse, error) {
	l := logic.NewDeleteSubCommentsByRooIDLogic(ctx, s.svcCtx)
	return l.DeleteSubCommentsByRooID(in)
}

func (s *CommentServer) DeleteAllCommentsBySubjectID(ctx context.Context, in *comment.DeleteAllCommentsBySubjectIDRequest) (*comment.DeleteAllCommentsBySubjectIDResponse, error) {
	l := logic.NewDeleteAllCommentsBySubjectIDLogic(ctx, s.svcCtx)
	return l.DeleteAllCommentsBySubjectID(in)
}

func (s *CommentServer) CreateSubject(ctx context.Context, in *comment.CreateSubjectRequest) (*comment.CreateSubjectResponse, error) {
	l := logic.NewCreateSubjectLogic(ctx, s.svcCtx)
	return l.CreateSubject(in)
}

func (s *CommentServer) GetSubject(ctx context.Context, in *comment.GetSubjectRequest) (*comment.GetSubjectResponse, error) {
	l := logic.NewGetSubjectLogic(ctx, s.svcCtx)
	return l.GetSubject(in)
}

func (s *CommentServer) DeleteSubject(ctx context.Context, in *comment.DeleteSubjectRequest) (*comment.DeleteSubjectResponse, error) {
	l := logic.NewDeleteSubjectLogic(ctx, s.svcCtx)
	return l.DeleteSubject(in)
}
