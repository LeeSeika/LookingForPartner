package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	WxUid        string `gorm:"size:128;primarykey"`
	Username     string `gorm:"128"`
	Avatar       string `gorm:"256"`
	School       string `gorm:"128"`
	Grade        int64
	Introduction string
	PostCount    int64

	// has many
	Posts    []Post    `gorm:"foreignKey:AuthorID"`
	Projects []Project `gorm:"foreignKey:MaintainerID"`

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
