package converter

import (
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/api/internal/types"
)

func CommentApiToRpc(commentInfo *types.Comment) comment.CommentInfo {
	return comment.CommentInfo{
		CommentID:        commentInfo.CommentID,
		SubjectID:        commentInfo.SubjectID,
		RootID:           commentInfo.RootID,
		ParentID:         commentInfo.ParentID,
		DialogID:         commentInfo.DialogID,
		AuthorID:         commentInfo.AuthorID,
		Content:          commentInfo.Content,
		Floor:            int32(commentInfo.Floor),
		CreatedAt:        commentInfo.CreatedAt,
		SubCommentsCount: int32(commentInfo.SubCommentCount),
	}
}

func CommentRpcToApi(commentInfo *comment.CommentInfo) types.Comment {

}
