package post

import "forum/internal/models"

func (s *PostService) GetAllPostsByCategory(categoryName string, userId int, filterBy string) ([]models.Post, error) {
	var err error
	var posts []models.Post

	switch filterBy {
	case "liked":
		posts, err = s.PostRepository.SELECT_postsByMostLiked(userId, categoryName)
		if err != nil {
			return nil, err
		}
	case "disliked":
		posts, err = s.PostRepository.SELECT_postsByMostDisliked(userId, categoryName)
		if err != nil {
			return nil, err
		}
	default:
		posts, err = s.PostRepository.SELECT_postsByCategory(userId, categoryName)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}
