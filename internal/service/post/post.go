package post

import (
	"forum/internal/repository/category"
	"forum/internal/repository/post"
)

type PostService struct {
	PostRepository     *post.PostRepository
	CategoryRepository *category.CategoryRepository
}

func NewPostService(repo *post.PostRepository, catRepo *category.CategoryRepository) *PostService {
	return &PostService{
		PostRepository:     repo,
		CategoryRepository: catRepo,
	}
}
