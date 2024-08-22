package converter

import (
	model2 "lookingforpartner/model"
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
		CreatedAt: poWithProj.Post.CreatedAt.Format(time.DateTime),
		Title:     poWithProj.Post.Title,
		Project:   projResp,
		Content:   poWithProj.Post.Content,
		AuthorID:  poWithProj.Post.AuthorID,
	}

	return &poInfo
}

func PostAndProject2PostInfo(po *model2.Post, proj *model2.Project) *post.PostInfo {
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

func Project2ProjResp(proj *model2.Project) *post.Project {
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
