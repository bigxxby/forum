package user

import (
	"forum/internal/repository/user"
)

type UserService struct {
	UserRepository *user.UserRepository
}

func NewUserService(repo *user.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}
