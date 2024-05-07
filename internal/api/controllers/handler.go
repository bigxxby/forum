package controllers

import (
	"forum/internal/service"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}
