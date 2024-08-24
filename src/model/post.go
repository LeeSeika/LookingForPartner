package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	PostID  string `gorm:"size:128;primarykey"`
	Title   string `gorm:"size:256"`
	Content string

	// has one
	Project Project `gorm:"foreignKey:ProjectID"`
	// belongs to
	AuthorID string `gorm:"size:128"`
	Author   User   `gorm:"foreignKey:AuthorID"`

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
