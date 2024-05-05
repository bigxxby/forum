package controllers

import (
	"forum/internal/database"
	"forum/internal/database/repos"
)

type Manager struct {
	Database *database.Database
	UserRepo *repos.UserRepo
}

func NewManager(Database database.Database) (*Manager, error) {
	userRepo := repos.NewUserRepo(Database.DB)
	manager := Manager{
		Database: &Database,
		UserRepo: userRepo,
	}

	return &manager, nil
}
