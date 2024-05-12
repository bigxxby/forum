package repository

import (
	"database/sql"
	"forum/internal/repository/category"
	"forum/internal/repository/post"
	"forum/internal/repository/user"
)

type Repository struct {
	UserRepository     *user.UserRepository
	CategoryRepository *category.CategoryRepository
	PostRepository     *post.PostRepository
}

func NewRepository(db *sql.DB) *Repository {
	// Инициализация репозиториев для пользователей, категорий и постов
	userRepo := user.NewUserRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	postRepo := post.NewPostRepository(db)

	// Создание экземпляра Repository с инициализированными репозиториями
	return &Repository{
		UserRepository:     userRepo,
		CategoryRepository: categoryRepo,
		PostRepository:     postRepo,
	}
}
