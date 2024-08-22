package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	WxUid        string `gorm:"primarykey"`
	Username     string
	Avatar       string
	School       string
	Grade        int64
	Introduction string
	PostCount    int64

	// has many
	Posts    []Post
	Projects []Project

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
