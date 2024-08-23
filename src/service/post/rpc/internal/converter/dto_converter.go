package converter

import (
	"time"

	"lookingforpartner/model"
	"lookingforpartner/service/post/rpc/pb/post"
)

func PostDBToRPC(po *model.Post) *post.PostInfo {
	proj := po.Project

	projRPC := &post.Project{
		ProjectID:     proj.ProjectID,
		MaintainerID:  proj.MaintainerID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}

	poInfo := post.PostInfo{
		PostID:    po.PostID,
		CreatedAt: po.CreatedAt.Format(time.DateTime),
		Title:     po.Title,
		Project:   projRPC,
		Content:   po.Content,
		AuthorID:  po.AuthorID,
	}

	return &poInfo
}

func ProjectDBToRPC(proj *model.Project) *post.Project {
	projRPC := post.Project{
		ProjectID:     proj.ProjectID,
		MaintainerID:  proj.MaintainerID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}

	return &projRPC
}
