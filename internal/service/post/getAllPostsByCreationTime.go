package post

import "forum/internal/models"

func (s *PostService) GetAllPostsByCreationTime(userId int) ([]models.Post, error) {
	posts, err := s.PostRepository.SELECT_postsCreatedAt(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
