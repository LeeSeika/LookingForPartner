package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	PostID  string `gorm:"primarykey"`
	Title   string
	Content string

	// has one
	Project Project `gorm:"foreignKey:ProjectID"`
	// belongs to
	AuthorID string
	Author   User `gorm:"foreignKey:AuthorID"`

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
