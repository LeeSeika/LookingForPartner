package dao

import (
	"lookingforpartner/common/dao"
	"lookingforpartner/model"
)

type PostInterface interface {
	CreatePost(post *model.Post) (*model.Post, error)
	DeletePost(postID string) (*model.Post, error)
	GetPost(postID string) (*model.Post, error)
	// todo: cursor
	GetPosts(page, size int64, order dao.OrderOpt) ([]*model.Post, error)
	GetPostsByAuthorID(page, size int64, authorID string, order dao.OrderOpt) ([]*model.Post, error)
	UpdateProject(project *model.Project) (*model.Project, error)
}
