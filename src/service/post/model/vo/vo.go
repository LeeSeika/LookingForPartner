package vo

import "lookingforpartner/service/post/model/entity"

type PostProject struct {
	*entity.Post
	*entity.Project
}
