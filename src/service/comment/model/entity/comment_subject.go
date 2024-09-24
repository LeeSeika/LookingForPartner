package entity

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectID        string `gorm:"size:128;index"`
	PostID           string `gorm:"size:128;uniqueIndex"`
	AllCommentCount  int    `gorm:"default:0"`
	RootCommentCount int    `gorm:"default:0"`
	Status           int8
}
