package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	WxUid        string `gorm:"index:idx_user_wxuid"`
	Username     string
	Avatar       string
	School       string
	Grade        int
	Introduction string
	PostCount    int
}
