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
