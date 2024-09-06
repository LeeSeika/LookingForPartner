package model

type CommentContent struct {
	CommentID string `gorm:"size:128;primaryKey"`
	Content   string
	MetaData  *string
}
