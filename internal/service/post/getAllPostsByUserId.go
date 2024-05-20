package post

import "forum/internal/models"

func (s *PostService) GetAllPostsByUserId(userId, userIdCreatedBy int) ([]models.Post, error) {

	posts, err := s.PostRepository.SELECT_createdByThisUser(userId, userIdCreatedBy)

	if err != nil {
		return nil, err
	}
	return posts, nil
}
