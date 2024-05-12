package post

import "forum/internal/models"

func (repo *PostRepository) SELECT_post(postId int) (*models.Post, error) {

	q := "SELECT * FROM posts WHERE id = ?"
	var post models.Post
	err := repo.DB.QueryRow(q, postId).Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil

}
