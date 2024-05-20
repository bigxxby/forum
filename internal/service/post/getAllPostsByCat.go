package post

import "forum/internal/models"

func (s *PostService) GetAllPostsByCategory(categoryName string, userId int) ([]models.Post, error) {

	posts, err := s.PostRepository.SELECT_postsByCategory(userId, categoryName)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
