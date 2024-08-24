package model

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ProjectID     string `gorm:"primarykey"`
	Name          string `gorm:"size:128"`
	Introduction  string
	Role          string `gorm:"size:40"`
	HeadCountInfo string
	Progress      string `gorm:"128"`

	// belongs to
	MaintainerID string `gorm:"size:128"`
	Maintainer   User   `gorm:"foreignKey:MaintainerID"`

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
