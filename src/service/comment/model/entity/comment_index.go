package entity

import (
	"gorm.io/gorm"
)

type CommentIndex struct {
	gorm.Model
	CommentID       string  `gorm:"size:128;index"`
	SubjectID       string  `gorm:"size:128;index"`
	RootID          *string `gorm:"size:128;index"`
	ParentID        *string `gorm:"size:128"`
	DialogID        *string `gorm:"size:128"`
	AuthorID        string  `gorm:"size:128;index"`
	SubCount        int     `gorm:"default:0"`
	LikeCount       int     `gorm:"default:0"`
	Floor           int
	SubCommentCount int `gorm:"default:0"`
	Status          int8
}
