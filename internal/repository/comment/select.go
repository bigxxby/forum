package comment

import (
	"forum/internal/models"
)

// замените на правильный путь к вашему пакету models

func (repo *CommentRepo) SELECT_Comments(postId int) ([]models.Comment, error) {
	q := `
    SELECT c.id, c.post_id, c.user_id, c.content, c.edited, u.login, c.created_at 
    FROM comments c 
    JOIN users u ON u.id = c.user_id
    WHERE c.post_id = ?
	ORDER BY created_at DESC
    `
	rows, err := repo.DB.Query(q, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.Edited,
			&comment.CreatedBy,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
