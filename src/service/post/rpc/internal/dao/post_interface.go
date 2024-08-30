package dao

import (
	"context"
	"lookingforpartner/service/post/model"

	basedao "lookingforpartner/common/dao"
)

type PostInterface interface {
	CreatePost(ctx context.Context, post *model.Post, proj *model.Project, idempotencyKey int64) (*model.PostProject, error)
	DeletePost(ctx context.Context, postID string, idempotencyKey int64) (*model.PostProject, error)
	GetPost(ctx context.Context, postID string) (*model.PostProject, error)
	GetPosts(ctx context.Context, page, size int64, order basedao.OrderOpt) ([]*model.PostProject, *basedao.Paginator, error)
	GetPostsByAuthorID(ctx context.Context, page, size int64, authorID string, order basedao.OrderOpt) ([]*model.PostProject, *basedao.Paginator, error)
	UpdateProject(ctx context.Context, project *model.Project) (*model.Project, error)
}
