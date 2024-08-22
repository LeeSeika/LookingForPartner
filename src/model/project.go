package model

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ProjectID     string `gorm:"primarykey"`
	Name          string
	Introduction  string
	Role          string
	HeadCountInfo string
	Progress      string

	// belongs to
	MaintainerID string
	Maintainer   User `gorm:"foreignKey:MaintainerID"`

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
