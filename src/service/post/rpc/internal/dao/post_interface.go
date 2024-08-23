package dao

import (
	"context"

	basedao "lookingforpartner/common/dao"
	"lookingforpartner/model"
)

type PostInterface interface {
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, error)
	DeletePost(ctx context.Context, postID string) (*model.Post, error)
	GetPost(ctx context.Context, postID string) (*model.Post, error)
	GetPosts(ctx context.Context, page, size int64, order basedao.OrderOpt) ([]*model.Post, *basedao.Paginator, error)
	GetPostsByAuthorID(ctx context.Context, page, size int64, authorID string, order basedao.OrderOpt) ([]*model.Post, *basedao.Paginator, error)
	UpdateProject(ctx context.Context, project *model.Project) (*model.Project, error)
}
