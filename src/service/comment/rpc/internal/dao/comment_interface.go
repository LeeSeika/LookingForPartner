package dao

import (
	"context"
	basedao "lookingforpartner/common/dao"
	"lookingforpartner/service/comment/model/entity"
	"lookingforpartner/service/comment/model/vo"
)

type CommentInterface interface {
	CreateComment(ctx context.Context, commentIndex *entity.CommentIndex, commentContent *entity.CommentContent) (*vo.CommentIndexContent, error)
	GetComment(ctx context.Context, commentID string) (*vo.CommentIndexContent, error)
	GetRootCommentsByPostID(ctx context.Context, postID string, page, size int64, order basedao.OrderOpt) ([]*vo.CommentIndexContent, *basedao.Paginator, error)
	GetTopSubCommentsByRootIDs(ctx context.Context, rootIDs []string, topCount int, order basedao.OrderOpt) ([]*vo.CommentIndexContent, error)
	DeleteComment(ctx context.Context, commentID string) (*vo.CommentIndexContent, error)
	DeleteSubCommentsByRootID(ctx context.Context, rootID string) error
	DeleteAllCommentsBySubjectID(ctx context.Context, subjectID string) error

	CreateSubject(ctx context.Context, subject *entity.Subject, idempotencyKey int64) (*entity.Subject, error)
	UpdateSubject(ctx context.Context, updatedSubject *entity.Subject) (*entity.Subject, error)
	DeleteSubject(ctx context.Context, subjectID string) (*entity.Subject, error)
	GetSubject(ctx context.Context, subjectID string) (*entity.Subject, error)
}
