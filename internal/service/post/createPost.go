package post

import "forum/internal/models"

func (s *PostService) CreatePost(userId int, title, content string, categories []string) (int, error) {
	catIds, err := s.CategoryRepository.SELECT_categoriesByName(categories)
	if err != nil {
		return -1, err
	}
	if len(catIds) != len(categories) {
		return -1, models.ErrBadRequest
	}

	postId, err := s.PostRepository.INSERT_post(userId, title, content, catIds)
	if err != nil {
		return -1, err
	}
	err = s.CategoryRepository.UPDATE_catCount(catIds)
	if err != nil {
		return -1, err
	}

	return postId, nil
}
