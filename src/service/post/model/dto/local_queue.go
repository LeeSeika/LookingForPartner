package dto

type UpdateUserPostCountMessage struct {
	Topic string
	Val   []byte
}

type DeleteSubjectMessage struct {
	Topic     string
	SubjectID string
}
