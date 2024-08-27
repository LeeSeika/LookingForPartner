package model

import (
	"gorm.io/gorm"
)

type Project struct {
	ProjectID     string `gorm:"size:128;index"`
	Name          string `gorm:"size:128"`
	Introduction  string
	Role          string `gorm:"size:40"`
	HeadCountInfo string
	Progress      string `gorm:"size:128"`

	// belongs to
	MaintainerID string `gorm:"size:128"`
	Maintainer   User   `gorm:"foreignKey:MaintainerID"`

	// base fields
	gorm.Model
}
