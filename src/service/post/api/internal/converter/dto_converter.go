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

	author := types.UserInfo{
		WxUid:        po.Author.WxUid,
		Avatar:       po.Author.Avatar,
		School:       po.Author.School,
		Grade:        po.Author.Grade,
		Introduction: po.Author.Introduction,
		PostCount:    po.Author.PostCount,
		Username:     po.Author.Username,
	}

	return types.Post{
		PostID:    po.PostID,
		CreatedAt: po.CreatedAt,
		Title:     po.Title,
		Project:   proj,
		Content:   po.Content,
		Author:    author,
	}
}
