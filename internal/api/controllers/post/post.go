package post

import (
	"forum/internal/service/post"
	"forum/internal/service/user"
)

type PostController struct {
	PostService *post.PostService
	UserService *user.UserService
}

func NewPostController(postService *post.PostService, userService *user.UserService) *PostController {
	return &PostController{
		PostService: postService,
		UserService: userService,
	}
}
