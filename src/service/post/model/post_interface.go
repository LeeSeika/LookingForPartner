package model

import "lookingforpartner/common/dao"

type PostInterface interface {
	CreatePostWithProjectTx(post *Post, project *Project) (*Post, *Project, error)
	CreatePost(post *Post) (*Post, error)
	DeletePostTx(postID int64) (*Post, *Project, error)
	GetPost(postID int64) (*PostWithProject, error)
	GetPosts(page, size int64, order dao.OrderOpt) ([]*PostWithProject, error)
	GetPostsByAuthorID(page, size int64, authorID string, order dao.OrderOpt) ([]*PostWithProject, error)
	SetProject(project *Project) (*Project, error)
}
