// Code generated by goctl. DO NOT EDIT.
// Source: post.proto

package postclient

import (
	"context"

	"lookingforpartner/service/post/rpc/pb/post"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePostRequest          = post.CreatePostRequest
	CreatePostResponse         = post.CreatePostResponse
	DeletePostRequest          = post.DeletePostRequest
	DeletePostResponse         = post.DeletePostResponse
	GetPostRequest             = post.GetPostRequest
	GetPostResponse            = post.GetPostResponse
	GetPostsByAuthorIDRequest  = post.GetPostsByAuthorIDRequest
	GetPostsByAuthorIDResponse = post.GetPostsByAuthorIDResponse
	GetPostsRequest            = post.GetPostsRequest
	GetPostsResponse           = post.GetPostsResponse
	PostInfo                   = post.PostInfo
	Project                    = post.Project
	UpdateProjectRequest       = post.UpdateProjectRequest
	UpdateProjectResponse      = post.UpdateProjectResponse

	Post interface {
		CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
		DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
		GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
		GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
		GetPostsByAuthorID(ctx context.Context, in *GetPostsByAuthorIDRequest, opts ...grpc.CallOption) (*GetPostsByAuthorIDResponse, error)
		UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
	}

	defaultPost struct {
		cli zrpc.Client
	}
)

func NewPost(cli zrpc.Client) Post {
	return &defaultPost{
		cli: cli,
	}
}

func (m *defaultPost) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	client := post.NewPostClient(m.cli.Conn())
	return client.CreatePost(ctx, in, opts...)
}

func (m *defaultPost) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	client := post.NewPostClient(m.cli.Conn())
	return client.DeletePost(ctx, in, opts...)
}

func (m *defaultPost) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	client := post.NewPostClient(m.cli.Conn())
	return client.GetPost(ctx, in, opts...)
}

func (m *defaultPost) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	client := post.NewPostClient(m.cli.Conn())
	return client.GetPosts(ctx, in, opts...)
}

func (m *defaultPost) GetPostsByAuthorID(ctx context.Context, in *GetPostsByAuthorIDRequest, opts ...grpc.CallOption) (*GetPostsByAuthorIDResponse, error) {
	client := post.NewPostClient(m.cli.Conn())
	return client.GetPostsByAuthorID(ctx, in, opts...)
}

func (m *defaultPost) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	client := post.NewPostClient(m.cli.Conn())
	return client.UpdateProject(ctx, in, opts...)
}