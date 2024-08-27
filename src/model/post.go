package model

import (
	"gorm.io/gorm"
)

type Post struct {
	PostID  string `gorm:"size:128;index"`
	Title   string `gorm:"size:256"`
	Content string

	// has one
	Project Project `gorm:"foreignKey:ProjectID"`
	// belongs to
	AuthorID string `gorm:"size:128"`
	Author   User   `gorm:"foreignKey:AuthorID"`

	// base fields
	gorm.Model
}
