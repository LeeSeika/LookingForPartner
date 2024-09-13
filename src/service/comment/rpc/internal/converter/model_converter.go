package converter

import (
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/model/entity"
	"time"
)

func SubjectDBToRPC(subject *entity.Subject) *comment.SubjectInfo {
	return &comment.SubjectInfo{
		SubjectID:        subject.SubjectID,
		PostID:           subject.PostID,
		AllCommentCount:  int32(subject.AllCommentCount),
		RootCommentCount: int32(subject.RootCommentCount),
		Status:           int32(subject.Status),
	}
}

func SingleCommentDBToRPC(commentIndex *entity.CommentIndex, commentContent *entity.CommentContent) *comment.CommentInfo {
	var rootID, parentID, dialogID string
	if commentIndex.RootID == nil {
		rootID = ""
	}
	if commentIndex.ParentID == nil {
		parentID = ""
	}
	if commentIndex.DialogID == nil {
		dialogID = ""
	}

	return &comment.CommentInfo{
		CommentID:        commentIndex.CommentID,
		SubjectID:        commentIndex.SubjectID,
		RootID:           rootID,
		ParentID:         parentID,
		DialogID:         dialogID,
		AuthorID:         commentIndex.AuthorID,
		Content:          commentContent.Content,
		Floor:            int32(commentIndex.Floor),
		CreatedAt:        commentIndex.CreatedAt.Format(time.DateTime),
		SubCommentsCount: int32(commentIndex.SubCommentCount),
		SubComments:      nil,
	}
}
