package post

import "forum/internal/models"

func (s *PostService) GetAllPosts(userId int, filter string) ([]models.Post, error) {
	var posts []models.Post
	var err error
	switch filter {
	case "liked":
		posts, err = s.PostRepository.SELECT_postsMostLiked(userId)
		if err != nil {
			return nil, err
		}
	case "disliked":
		posts, err = s.PostRepository.SELECT_postsMostDisliked(userId)
		if err != nil {
			return nil, err
		}
	default:
		posts, err = s.PostRepository.SELECT_postsCreatedAt(userId)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}
