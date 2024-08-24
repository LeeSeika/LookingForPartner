package converter

import (
	"lookingforpartner/pb/post"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/api/internal/types"
)

func ProjectRpcToApi(proj *post.Project) types.Project {
	maintainer := types.UserInfo{
		WxUid:        proj.Maintainer.WxUid,
		Avatar:       proj.Maintainer.Avatar,
		School:       proj.Maintainer.School,
		Grade:        proj.Maintainer.Grade,
		Introduction: proj.Maintainer.Introduction,
		PostCount:    proj.Maintainer.PostCount,
		Username:     proj.Maintainer.Username,
	}

	return types.Project{
		ProjectID:     proj.ProjectID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Maintainer:    maintainer,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}
}

func ProjectApiToRpc(proj *types.Project) post.Project {
	maintainer := &user.UserInfo{
		WxUid:        proj.Maintainer.WxUid,
		Avatar:       proj.Maintainer.Avatar,
		School:       proj.Maintainer.School,
		Grade:        proj.Maintainer.Grade,
		Introduction: proj.Maintainer.Introduction,
		PostCount:    proj.Maintainer.PostCount,
		Username:     proj.Maintainer.Username,
	}

	return post.Project{
		ProjectID:     proj.ProjectID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Maintainer:    maintainer,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}
}

func PostRpcToApi(po *post.PostInfo) types.Post {
	proj := types.Project{}
	if po.GetProject() != nil {
		proj = ProjectRpcToApi(po.GetProject())
	}

	return types.Post{
		PostID:    po.PostID,
		CreatedAt: po.CreatedAt,
		Title:     po.Title,
		Project:   proj,
		Content:   po.Content,
		AuthorID:  po.AuthorID,
	}
}
