package user

import (
	"forum/internal/service/user"
)

type UserController struct {
	UserService *user.UserService
}

func NewUserController(service *user.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}
