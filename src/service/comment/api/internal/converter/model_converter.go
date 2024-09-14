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

	commentApi := commentRpcToApi(commentInfo)

	subCommentApis := make([]*types.Comment, 0, commentInfo.SubCommentsCount)

	for i := 0; i < int(commentInfo.SubCommentsCount); i++ {
		subCommentApi := commentRpcToApi(commentInfo.SubComments[i])
		subCommentApis = append(subCommentApis, &subCommentApi)
	}

	commentApi.SubComments = subCommentApis

	return commentApi
}

func commentRpcToApi(commentRpc *comment.CommentInfo) types.Comment {
	return types.Comment{
		CommentID:       commentRpc.CommentID,
		SubjectID:       commentRpc.SubjectID,
		RootID:          commentRpc.RootID,
		ParentID:        commentRpc.ParentID,
		DialogID:        commentRpc.DialogID,
		AuthorID:        commentRpc.AuthorID,
		LikeCount:       0,
		Floor:           int(commentRpc.Floor),
		CreatedAt:       commentRpc.CreatedAt,
		SubCommentCount: int(commentRpc.SubCommentsCount),
		Content:         commentRpc.Content,
	}
}
