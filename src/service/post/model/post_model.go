package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostID   int64 `gorm:"index:idx_post_id,unique"`
	AuthorID string
	Title    string
	Content  string
}

type Project struct {
	gorm.Model
	ProjectID     int64 `gorm:"index:idx_project_id,unique"`
	MaintainerID  string
	Maintainer    string
	Name          string
	Introduction  string
	Role          string
	HeadCountInfo string
	Progress      string
	PostID        int64
}

// dto

type PostWithProject struct {
	Post
	Project
}
