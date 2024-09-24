package dao

import (
	"context"
	"lookingforpartner/service/post/model/entity"
	"lookingforpartner/service/post/model/vo"

	basedao "lookingforpartner/common/dao"
)

type PostInterface interface {
	CreatePost(ctx context.Context, post *entity.Post, proj *entity.Project, idempotencyKey int64) (*vo.PostProject, error)
	DeletePost(ctx context.Context, postID string) (*vo.PostProject, error)
	UpdatePost(ctx context.Context, updatedPost *entity.Post) (*entity.Post, error)
	GetPost(ctx context.Context, postID string) (*vo.PostProject, error)
	GetPosts(ctx context.Context, page, size int64, order basedao.OrderOpt) ([]*vo.PostProject, *basedao.Paginator, error)
	GetPostsByAuthorID(ctx context.Context, page, size int64, authorID string, order basedao.OrderOpt) ([]*vo.PostProject, *basedao.Paginator, error)
	UpdateProject(ctx context.Context, project *entity.Project) (*entity.Project, error)
}
