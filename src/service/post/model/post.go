package model

import (
	"gorm.io/gorm"
)

type Post struct {
	PostID   string `gorm:"size:128;index"`
	Title    string `gorm:"size:256"`
	Content  string
	AuthorID string `gorm:"size:128"`

	// base fields
	gorm.Model
}

type PostProject struct {
	*Post
	*Project
}
