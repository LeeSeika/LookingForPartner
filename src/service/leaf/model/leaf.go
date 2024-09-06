package model

import (
	"time"

	"gorm.io/gorm"
)

type Leaf struct {
	BizTag string `gorm:"primarykey"`
	MaxId  int64  `gorm:"not null;default:0"`

	// base fields
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
