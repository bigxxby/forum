package post

import "forum/internal/models"

func (p *PostService) GetAllLikedPosts(userId int) ([]models.Post, error) {
	posts, err := p.PostRepository.SELECT_liked_posts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
