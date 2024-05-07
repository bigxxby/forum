package post

import "forum/internal/repository/post"

type PostService struct {
	PostRepository *post.PostRepository
}

func NewPostService(repo *post.PostRepository) *PostService {
	return &PostService{
		PostRepository: repo,
	}
}
