package converter

import (
	"lookingforpartner/pb/post"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/post/api/internal/types"
)

func ProjectRpcToApi(projectRpc *post.Project) types.Project {
	maintainer := types.UserInfo{
		WxUid:        projectRpc.Maintainer.WxUid,
		Avatar:       projectRpc.Maintainer.Avatar,
		School:       projectRpc.Maintainer.School,
		Grade:        projectRpc.Maintainer.Grade,
		Introduction: projectRpc.Maintainer.Introduction,
		PostCount:    projectRpc.Maintainer.PostCount,
		Username:     projectRpc.Maintainer.Username,
	}

	return types.Project{
		ProjectID:     projectRpc.ProjectID,
		Name:          projectRpc.Name,
		Introduction:  projectRpc.Introduction,
		Maintainer:    maintainer,
		Role:          projectRpc.Role,
		HeadCountInfo: projectRpc.HeadCountInfo,
		Progress:      projectRpc.Progress,
	}
}

func ProjectApiToRpc(projectApi *types.Project) post.Project {
	maintainer := &user.UserInfo{
		WxUid:        projectApi.Maintainer.WxUid,
		Avatar:       projectApi.Maintainer.Avatar,
		School:       projectApi.Maintainer.School,
		Grade:        projectApi.Maintainer.Grade,
		Introduction: projectApi.Maintainer.Introduction,
		PostCount:    projectApi.Maintainer.PostCount,
		Username:     projectApi.Maintainer.Username,
	}

	return post.Project{
		ProjectID:     projectApi.ProjectID,
		Name:          projectApi.Name,
		Introduction:  projectApi.Introduction,
		Maintainer:    maintainer,
		Role:          projectApi.Role,
		HeadCountInfo: projectApi.HeadCountInfo,
		Progress:      projectApi.Progress,
	}
}

func PostRpcToApi(postRpc *post.PostInfo) types.Post {
	proj := types.Project{}
	if postRpc.GetProject() != nil {
		proj = ProjectRpcToApi(postRpc.GetProject())
	}

	author := types.UserInfo{
		WxUid:        postRpc.Author.WxUid,
		Avatar:       postRpc.Author.Avatar,
		School:       postRpc.Author.School,
		Grade:        postRpc.Author.Grade,
		Introduction: postRpc.Author.Introduction,
		PostCount:    postRpc.Author.PostCount,
		Username:     postRpc.Author.Username,
	}

	return types.Post{
		PostID:    postRpc.PostID,
		CreatedAt: postRpc.CreatedAt,
		Title:     postRpc.Title,
		Project:   proj,
		Content:   postRpc.Content,
		Author:    author,
	}
}
