package controllers

import (
	categoryController "forum/internal/api/controllers/category"
	commentController "forum/internal/api/controllers/comment"
	postController "forum/internal/api/controllers/post"
	userController "forum/internal/api/controllers/user"
	"forum/internal/service"
)

type Controller struct {
	UserController     *userController.UserController
	CategoryController *categoryController.CategoryController
	PostController     *postController.PostController
	HTMLController     *HTMLController
	CommentController  *commentController.CommentController
}

func NewController(service *service.Service) *Controller {
	userControllerVar := userController.NewUserController(service.User)
	categoryControllerVar := categoryController.NewCategoryController(service.Category)
	postControllerVar := postController.NewPostController(service.Post, service.User, service.Category)
	commentControllerVar := commentController.NewCommentController(service.Comment)
	return &Controller{
		UserController:     userControllerVar,
		CategoryController: categoryControllerVar,
		PostController:     postControllerVar,
		CommentController:  commentControllerVar,
	}
}
