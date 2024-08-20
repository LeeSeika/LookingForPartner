package converter

import (
	"lookingforpartner/service/post/api/internal/types"
	"lookingforpartner/service/post/rpc/pb/post"
)

func ProjectRpc2Api(proj *post.Project) types.Project {
	return types.Project{
		ProjectID:     proj.ProjectID,
		MaintainerID:  proj.MaintainerID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Maintainer:    proj.Maintainer,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}
}

func ProjectApi2Rpc(proj *types.Project) post.Project {
	return post.Project{
		ProjectID:     proj.ProjectID,
		MaintainerID:  proj.MaintainerID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Maintainer:    proj.Maintainer,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}
}

func PostRpc2Api(po *post.PostInfo) types.Post {
	proj := types.Project{}
	if po.GetProject() != nil {
		proj = ProjectRpc2Api(po.GetProject())
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
