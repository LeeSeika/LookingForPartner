package model

import (
	"gorm.io/gorm"
)

type User struct {
	WxUid        string `gorm:"size:128;index"`
	Username     string `gorm:"size:128"`
	Avatar       string `gorm:"size:256"`
	School       string `gorm:"size:128"`
	Grade        int64
	Introduction string
	PostCount    int64

	// base fields
	gorm.Model
}
