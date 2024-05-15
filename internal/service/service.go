package service

import (
	"forum/internal/repository"
	"forum/internal/service/category"
	"forum/internal/service/comment"
	"forum/internal/service/post"
	"forum/internal/service/user"
)

type Service struct {
	Category *category.CategoryService
	User     *user.UserService
	Post     *post.PostService
	Comment  *comment.CommentService
}

func NewService(repo *repository.Repository) *Service {
	userService := user.NewUserService(repo.UserRepository)
	categoryService := category.NewCategoryService(repo.CategoryRepository)
	postService := post.NewPostService(repo.PostRepository, categoryService.CategoryRepository)
	commentService := comment.NewCommentService(repo.CommentRepo)

	return &Service{
		Category: categoryService,
		User:     userService,
		Post:     postService,
		Comment:  commentService,
	}
}
