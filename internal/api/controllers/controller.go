package controllers

import (
	categoryController "forum/internal/api/controllers/category"
	postController "forum/internal/api/controllers/post"
	userController "forum/internal/api/controllers/user"
	"forum/internal/service"
)

type Controller struct {
	UserController     *userController.UserController
	CategoryController *categoryController.CategoryController
	PostController     *postController.PostController
	HTMLController     *HTMLController
}

func NewController(service *service.Service) *Controller {
	userControllerVar := userController.NewUserController(service.User)
	categoryControllerVar := categoryController.NewCategoryController(service.Category)
	postControllerVar := postController.NewPostController(service.Post, service.User, service.Category)
	return &Controller{
		UserController:     userControllerVar,
		CategoryController: categoryControllerVar,
		PostController:     postControllerVar,
	}
}
