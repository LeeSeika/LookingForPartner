package converter

import (
	"lookingforpartner/pb/user"
	"time"

	"lookingforpartner/model"
	"lookingforpartner/pb/post"
)

func PostDBToRPC(po *model.Post) *post.PostInfo {
	proj := po.Project

	maintainer := &user.UserInfo{
		WxUid:        proj.Maintainer.WxUid,
		Username:     proj.Maintainer.Username,
		Avatar:       proj.Maintainer.Avatar,
		School:       proj.Maintainer.School,
		Grade:        proj.Maintainer.Grade,
		Introduction: proj.Maintainer.Introduction,
		PostCount:    proj.Maintainer.PostCount,
	}

	projRPC := &post.Project{
		ProjectID:     proj.ProjectID,
		Maintainer:    maintainer,
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
	maintainer := &user.UserInfo{
		WxUid:        proj.Maintainer.WxUid,
		Username:     proj.Maintainer.Username,
		Avatar:       proj.Maintainer.Avatar,
		School:       proj.Maintainer.School,
		Grade:        proj.Maintainer.Grade,
		Introduction: proj.Maintainer.Introduction,
		PostCount:    proj.Maintainer.PostCount,
	}

	projRPC := post.Project{
		ProjectID:     proj.ProjectID,
		Maintainer:    maintainer,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}

	return &projRPC
}
