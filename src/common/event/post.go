package event

import (
	"lookingforpartner/service/post/model/entity"
)

type CreatePost struct {
	IdempotencyKey int64
	Post           *entity.Post
	Project        *entity.Project
}

type DeletePost struct {
	Post    *entity.Post
	Project *entity.Project
}
