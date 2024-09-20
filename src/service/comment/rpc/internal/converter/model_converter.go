package converter

import (
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/model/entity"
	"time"
)

func SubjectDBToRPC(subjectDB *entity.Subject) *comment.SubjectInfo {
	return &comment.SubjectInfo{
		SubjectID:        subjectDB.SubjectID,
		PostID:           subjectDB.PostID,
		AllCommentCount:  int32(subjectDB.AllCommentCount),
		RootCommentCount: int32(subjectDB.RootCommentCount),
		Status:           int32(subjectDB.Status),
	}
}

func SingleCommentDBToRPC(commentIndexRpc *entity.CommentIndex, commentContentRpc *entity.CommentContent) *comment.CommentInfo {
	var rootID, parentID, dialogID string
	if commentIndexRpc.RootID == nil {
		rootID = ""
	}
	if commentIndexRpc.ParentID == nil {
		parentID = ""
	}
	if commentIndexRpc.DialogID == nil {
		dialogID = ""
	}

	return &comment.CommentInfo{
		CommentID:        commentIndexRpc.CommentID,
		SubjectID:        commentIndexRpc.SubjectID,
		RootID:           rootID,
		ParentID:         parentID,
		DialogID:         dialogID,
		AuthorID:         commentIndexRpc.AuthorID,
		Content:          commentContentRpc.Content,
		Floor:            int32(commentIndexRpc.Floor),
		CreatedAt:        commentIndexRpc.CreatedAt.Format(time.DateTime),
		SubCommentsCount: int32(commentIndexRpc.SubCommentCount),
		SubComments:      nil,
	}
}
