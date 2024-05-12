package post

import (
	"forum/internal/service/category"
	"forum/internal/service/post"
	"forum/internal/service/user"
)

type PostController struct {
	PostService     *post.PostService
	UserService     *user.UserService
	CategoryService *category.CategoryService
}

func NewPostController(postService *post.PostService, userService *user.UserService, categoryService *category.CategoryService) *PostController {
	return &PostController{
		PostService:     postService,
		UserService:     userService,
		CategoryService: categoryService,
	}
}
