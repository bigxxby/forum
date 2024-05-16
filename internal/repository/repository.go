package repository

import (
	"database/sql"
	"forum/internal/repository/category"
	"forum/internal/repository/comment"
	"forum/internal/repository/likes"
	"forum/internal/repository/post"
	"forum/internal/repository/user"
)

type Repository struct {
	UserRepository     *user.UserRepository
	CategoryRepository *category.CategoryRepository
	PostRepository     *post.PostRepository
	CommentRepo        *comment.CommentRepo
	LikesRepo          *likes.LikesRepo
}

func NewRepository(db *sql.DB) *Repository {
	userRepo := user.NewUserRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	postRepo := post.NewPostRepository(db)
	commentRepo := comment.NewCommentRepo(db)
	likesRepo := likes.NewLikesRepo(db)

	return &Repository{
		UserRepository:     userRepo,
		CategoryRepository: categoryRepo,
		PostRepository:     postRepo,
		CommentRepo:        commentRepo,
		LikesRepo:          likesRepo,
	}
}
