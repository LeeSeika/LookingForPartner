package entity

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
	MaintainerID  string `gorm:"size:128;index"`
	PostID        string `gorm:"size:128;index"`

	// base fields
	gorm.Model
}
