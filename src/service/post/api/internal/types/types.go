// Code generated by goctl. DO NOT EDIT.
package types

type CreatePostRequest struct {
	Title   string  `json:"title"`
	Project Project `json:"project"`
	Content string  `json:"content"`
}

type CreatePostResponse struct {
	PostID    int64   `json:"post_id, string"`
	CreatedAt string  `json:"created_at"`
	Title     string  `json:"title"`
	Project   Project `json:"project"`
	Content   string  `json:"content"`
}

type DeletePostRequest struct {
	PostID int64 `json:"post_id, string"`
}

type DeletePostResponse struct {
}

type GetPostByAuthorIDRequest struct {
	AuthorID string `json:"author_id"`
	Page     int64  `json:"page"`
	Size     int64  `json:"size"`
	Order    string `json:"order"`
}

type GetPostRequest struct {
	PostID int64 `json:"post_id, string"`
}

type GetPostResponse struct {
	Post Post `json:"post"`
}

type GetPostsByAuthorIDResponse struct {
	Posts []GetPostResponse `json:"posts"`
}

type GetPostsRequest struct {
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
	Order string `json:"order"`
}

type GetPostsResponse struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	PostID    int64   `json:"post_id, string"`
	CreatedAt string  `json:"created_at"`
	Title     string  `json:"title"`
	Project   Project `json:"project"`
	Content   string  `json:"content"`
	AuthorID  string  `json:"author_id"`
}

type Project struct {
	ProjectID     int64  `json:"project_id, string"`
	MaintainerID  string `json:"maintainer_id"`
	Name          string `json:"name"`
	Introduction  string `json:"introduction"`
	Maintainer    string `json:"maintainer"`
	Role          string `json:"role"`
	HeadCountInfo string `json:"head_count_info"`
	Progress      string `json:"progress"`
}

type UpdateProjectRequest struct {
	Name          string `json:"name"`
	Introduction  string `json:"introduction"`
	Role          string `json:"role"`
	HeadCountInfo string `json:"head_count_info"`
	Progress      string `json:"progress"`
}

type UpdateProjectResponse struct {
	Project Project `json:"project"`
}