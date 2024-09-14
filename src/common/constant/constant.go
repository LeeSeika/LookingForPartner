package constant

const (
	// NanoidPrefixIdempotence is a nanoid prefix for each process to implement idempotence
	NanoidPrefixIdempotence = "idem_"

	// NanoidPrefixUser is a nanoid prefix for each user
	NanoidPrefixUser = "wx_"

	// NanoidPrefixPost is a nanoid prefix for each post
	NanoidPrefixPost = "po_"

	// NanoidPrefixProject is a nanoid prefix for each project
	NanoidPrefixProject = "proj_"

	// NanoidPrefixComment is a nanoid prefix for each comment
	NanoidPrefixComment = "cmt_"

	// NanoidPrefixSubject is a nanoid prefix for each subject
	NanoidPrefixSubject = "subj_"
)

const (
	// MqMessageKeyDeleteAllCommentsBySubjectID is a key represents the task DeleteAllCommentsBySubjectID
	MqMessageKeyDeleteAllCommentsBySubjectID = "mq_delete-all-comments-by-subject-id"

	// MqMessageKeyDeleteSubCommentsByRootID is a key represents the task DeleteSubCommentsByRootID
	MqMessageKeyDeleteSubCommentsByRootID = "mq_delete-sub-comments-by-root-id"
)
