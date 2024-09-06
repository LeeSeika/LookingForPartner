package converter

import (
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/model"
	"time"

	"lookingforpartner/pb/post"
)

func PostDBToRPC(po *model.Post) *post.PostInfo {
	author := user.UserInfo{WxUid: po.AuthorID}
	poInfo := post.PostInfo{
		PostID:    po.PostID,
		CreatedAt: po.CreatedAt.Format(time.DateTime),
		Title:     po.Title,
		Content:   po.Content,
		Author:    &author,
	}

	return &poInfo
}

func ProjectDBToRPC(proj *model.Project) *post.Project {
	maintainer := user.UserInfo{WxUid: proj.MaintainerID}
	projRPC := post.Project{
		ProjectID:     proj.ProjectID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
		Maintainer:    &maintainer,
	}

	return &projRPC
}
