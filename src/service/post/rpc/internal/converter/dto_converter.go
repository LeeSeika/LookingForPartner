package converter

import (
	"lookingforpartner/service/post/model"
	"time"

	"lookingforpartner/pb/post"
)

func PostDBToRPC(po *model.Post) *post.PostInfo {
	poInfo := post.PostInfo{
		PostID:    po.PostID,
		CreatedAt: po.CreatedAt.Format(time.DateTime),
		Title:     po.Title,
		Content:   po.Content,
		AuthorID:  po.AuthorID,
	}

	return &poInfo
}

func ProjectDBToRPC(proj *model.Project) *post.Project {

	projRPC := post.Project{
		ProjectID:     proj.ProjectID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}

	return &projRPC
}
