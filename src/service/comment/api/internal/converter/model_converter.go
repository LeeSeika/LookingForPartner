package converter

import (
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/api/internal/types"
)

func CommentApiToRpc(commentApi *types.Comment) comment.CommentInfo {
	return comment.CommentInfo{
		CommentID:        commentApi.CommentID,
		SubjectID:        commentApi.SubjectID,
		RootID:           commentApi.RootID,
		ParentID:         commentApi.ParentID,
		DialogID:         commentApi.DialogID,
		AuthorID:         commentApi.AuthorID,
		Content:          commentApi.Content,
		Floor:            int32(commentApi.Floor),
		CreatedAt:        commentApi.CreatedAt,
		SubCommentsCount: int32(commentApi.SubCommentCount),
	}
}

func CommentRpcToApi(commentRpc *comment.CommentInfo) types.Comment {

	commentApi := commentRpcToApi(commentRpc)

	subCommentApis := make([]*types.Comment, 0, commentRpc.SubCommentsCount)

	for i := 0; i < int(commentRpc.SubCommentsCount); i++ {
		subCommentApi := commentRpcToApi(commentRpc.SubComments[i])
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
