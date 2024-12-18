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
	// DefaultLocalQueueChanCap is a default channel capacity for local queue
	DefaultLocalQueueChanCap = 100

	// DefaultLocalQueueDataCap is a default data slice capacity for local queue
	DefaultLocalQueueDataCap = 1000
)
