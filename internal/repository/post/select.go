package post

import (
	"forum/internal/models"
)

func (repo *PostRepository) SELECT_post(postId, userId int) (*models.Post, bool, error) {
	q := `
        SELECT p.id, p.user_id, p.title, p.content, p.created_at, u.login,
               p.likes, c.name,
               CASE WHEN l.post_id IS NOT NULL THEN true ELSE false END AS liked
        FROM posts p
        LEFT JOIN users u ON p.user_id = u.id
        LEFT JOIN categories c ON p.category_id = c.id
        LEFT JOIN likes l ON p.id = l.post_id AND l.user_id = ?
        WHERE p.id = ?;
    `
	var post models.Post
	var liked bool
	err := repo.DB.QueryRow(q, userId, postId).Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.CreatedBy, &post.Likes, &post.Category, &liked)
	if err != nil {
		return nil, false, err
	}
	return &post, liked, nil
}

func (repo *PostRepository) SELECT_postsCreatedAt(userId int) ([]models.Post, error) {
	q := `
        SELECT p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, c.name,
               CASE WHEN l.post_id IS NOT NULL THEN true ELSE false END AS liked
        FROM posts p
        LEFT JOIN users u ON p.user_id = u.id
        LEFT JOIN categories c ON p.category_id = c.id
        LEFT JOIN likes l ON p.id = l.post_id AND l.user_id = ?
        ORDER BY p.created_at DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post

		err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.CreatedBy, &post.Likes, &post.Category, &post.Liked)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
