package post

import "forum/internal/models"

func (repo *PostRepository) SELECT_post(postId int) (*models.Post, error) {

	q := `
		SELECT p.*, u.login
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.id = ?;
		`
	var post models.Post
	err := repo.DB.QueryRow(q, postId).Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.CreatedBy)
	if err != nil {
		return nil, err
	}
	return &post, nil

}
func (repo *PostRepository) SELECT_postsCreatedAt() ([]models.Post, error) {

	q := `
		SELECT p.*, u.login
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY created_at DESC;

	`

	var posts []models.Post

	rows, err := repo.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post

		err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.CreatedBy)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
