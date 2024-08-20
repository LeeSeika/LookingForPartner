package converter

import (
	"lookingforpartner/service/post/model"
	"lookingforpartner/service/post/rpc/pb/post"
	"time"
)

func PostWithProject2PostInfo(poWithProj *model.PostWithProject) *post.PostInfo {
	proj := poWithProj.Project

	projResp := &post.Project{
		ProjectID:     proj.ProjectID,
		MaintainerID:  proj.MaintainerID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Maintainer:    proj.Maintainer,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}

	poInfo := post.PostInfo{
		PostID:    poWithProj.Post.PostID,
		CreatedAt: poWithProj.CreatedAt.Format(time.DateTime),
		Title:     poWithProj.Title,
		Project:   projResp,
		Content:   poWithProj.Content,
		AuthorID:  poWithProj.AuthorID,
	}

	return &poInfo
}

func PostAndProject2PostInfo(po *model.Post, proj *model.Project) *post.PostInfo {
	var projResp *post.Project
	if proj != nil {
		projResp = &post.Project{
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

	poInfo := post.PostInfo{
		PostID:    po.PostID,
		CreatedAt: po.CreatedAt.Format(time.DateTime),
		Title:     po.Title,
		Project:   projResp,
		Content:   po.Content,
		AuthorID:  po.AuthorID,
	}

	return &poInfo
}

func Project2ProjResp(proj *model.Project) *post.Project {
	porjResp := post.Project{
		ProjectID:     proj.ProjectID,
		MaintainerID:  proj.MaintainerID,
		Name:          proj.Name,
		Introduction:  proj.Introduction,
		Maintainer:    proj.Maintainer,
		Role:          proj.Role,
		HeadCountInfo: proj.HeadCountInfo,
		Progress:      proj.Progress,
	}

	return &porjResp
}