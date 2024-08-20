// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"lookingforpartner/service/post/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/posts",
				Handler: GetPostsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/posts/:postID",
				Handler: GetPostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/posts/author/:authorID",
				Handler: GetPostsByAuthorIDHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/posts",
				Handler: CreatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/posts/:postID",
				Handler: DeletePostHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/posts/project/:projectID",
				Handler: UpdateProjectHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
