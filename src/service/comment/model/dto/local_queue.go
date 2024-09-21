package dto

type DeleteCommentMessage struct {
	Topic string
	Key   string
	Val   string
}

type UpdateUserPostCountMessage struct {
	Topic string
	Val   []byte
}
