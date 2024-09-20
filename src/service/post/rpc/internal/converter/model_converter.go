package converter

import (
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/model/entity"
	"time"

	"lookingforpartner/pb/post"
)

func PostDBToRPC(postDB *entity.Post) *post.PostInfo {
	author := user.UserInfo{WxUid: postDB.AuthorID}
	poRpc := post.PostInfo{
		PostID:    postDB.PostID,
		CreatedAt: postDB.CreatedAt.Format(time.DateTime),
		Title:     postDB.Title,
		Content:   postDB.Content,
		Author:    &author,
	}

	return &poRpc
}

func ProjectDBToRPC(projectDB *entity.Project) *post.Project {
	maintainer := user.UserInfo{WxUid: projectDB.MaintainerID}
	projRpc := post.Project{
		ProjectID:     projectDB.ProjectID,
		Name:          projectDB.Name,
		Introduction:  projectDB.Introduction,
		Role:          projectDB.Role,
		HeadCountInfo: projectDB.HeadCountInfo,
		Progress:      projectDB.Progress,
		Maintainer:    &maintainer,
	}

	return &projRpc
}
