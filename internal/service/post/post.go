package post

import (
	"forum/internal/repository/category"
	"forum/internal/repository/likes"
	"forum/internal/repository/post"
)

type PostService struct {
	PostRepository     *post.PostRepository
	CategoryRepository *category.CategoryRepository
	LikesRepo          *likes.LikesRepo
}

func NewPostService(repo *post.PostRepository, catRepo *category.CategoryRepository, likeRepo *likes.LikesRepo) *PostService {
	return &PostService{
		PostRepository:     repo,
		CategoryRepository: catRepo,
		LikesRepo:          likeRepo,
	}
}
