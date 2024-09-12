package vo

import "lookingforpartner/service/comment/model/entity"

type CommentIndexContent struct {
	*entity.CommentIndex
	*entity.CommentContent
}

type RootCommentWithSubs struct {
	*CommentIndexContent
	Sub []*CommentIndexContent
}
