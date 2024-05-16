package post

import (
	"forum/internal/models"
)

func (s *PostService) GetPostById(postId int, userId int) (*models.Post, error) {

	post, liked, err := s.PostRepository.SELECT_post(postId, userId)
	if err != nil {
		return nil, err
	}
	if liked {
		post.Liked = true
	}

	return post, nil
}
