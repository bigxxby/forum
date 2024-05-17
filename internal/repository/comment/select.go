package comment

import (
	"forum/internal/models"
)

func (repo *CommentRepo) SELECT_Comments(postId int, userId int) ([]models.Comment, error) {
	q := `
    SELECT  c.id, c.post_id , c.parent_id, c.user_id, c.content, c.edited, u.login, c.likes, c.created_at , CASE WHEN l.user_id IS NOT NULL THEN true ELSE false END 
    FROM comments c 
    LEFT JOIN users u ON u.id = c.user_id
    LEFT JOIN likes l ON l.user_id = ? AND c.id = l.comment_id
    WHERE c.post_id = ?
	ORDER BY created_at DESC
    `
	rows, err := repo.DB.Query(q, userId, postId)
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
			&comment.ParentId,
			&comment.UserID,
			&comment.Content,
			&comment.Edited,
			&comment.CreatedBy,
			&comment.Likes,
			&comment.CreatedAt,
			&comment.Liked,
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
func (repo *CommentRepo) SELECT_liked_comments(userId int) ([]models.Comment, error) {
	q := `
    SELECT c.id,  c.parent_id, c.post_id, c.user_id, c.content, c.edited, u.login, c.likes, c.created_at
    FROM comments c 
    JOIN users u ON u.id = c.user_id
	JOIN likes l ON l.user_id = ?
	WHERE l.comment_id = c.id
	ORDER BY l.id
    `
	rows, err := repo.DB.Query(q, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.ParentId,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.Edited,
			&comment.CreatedBy,
			&comment.Likes,
			&comment.CreatedAt,
		)
		comment.Liked = true
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
