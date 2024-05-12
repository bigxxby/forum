package post

import (
	"forum/internal/models"
)

func (s *PostService) GetPostById(postId int) (*models.Post, error) {

	post, err := s.PostRepository.SELECT_post(postId)
	if err != nil {
		return nil, err
	}
	return post, nil
}
