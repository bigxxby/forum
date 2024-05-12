package post

import "forum/internal/models"

func (s *PostService) GetAllPostsByCreationTime() ([]models.Post, error) {
	posts, err := s.PostRepository.SELECT_postsCreatedAt()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
