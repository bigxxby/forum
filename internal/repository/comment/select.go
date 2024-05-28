package comment

import (
	"forum/internal/models"
)

func (repo *CommentRepo) SELECT_Comments(postId int, userId int) ([]models.Comment, error) {
	q := `
    SELECT  
        c.id, 
        c.post_id, 
        c.user_id, 
        c.content, 
        c.edited, 
        u.login, 
        c.likes, 
        c.dislikes, 
        c.created_at, 
        CASE WHEN l1.user_id IS NOT NULL THEN true ELSE false END AS liked,
        CASE WHEN l2.user_id IS NOT NULL THEN true ELSE false END AS disliked
    FROM 
        comments c 
    LEFT JOIN 
        users u ON u.id = c.user_id
    LEFT JOIN 
        likes_dislikes l1 ON l1.user_id = ? AND c.id = l1.comment_id AND l1.value = true -- Лайки
	LEFT JOIN 
		likes_dislikes l2 ON l2.user_id = ? AND c.id = l2.comment_id AND l2.value = false -- дизы
	WHERE 
	c.post_id = ?
    ORDER BY 
        c.created_at DESC
    `
	rows, err := repo.DB.Query(q, userId, userId, postId)
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
			&comment.Likes,
			&comment.Dislikes,
			&comment.CreatedAt,
			&comment.Liked,
			&comment.Disliked,
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
    SELECT 
        c.id, 
        c.post_id, 
        c.user_id, 
        c.content, 
        c.edited, 
        u.login, 
        c.likes, 
        c.created_at,
        COALESCE(ld.likes_count, 0) AS dislikes, 
        CASE WHEN l.value = true THEN true ELSE false END AS liked,
        CASE WHEN d.value = false THEN true ELSE false END AS disliked
    FROM 
        comments c 
    JOIN 
        users u ON u.id = c.user_id
    JOIN 
        likes_dislikes l ON l.user_id = ? AND l.comment_id = c.id AND l.value = true 
    LEFT JOIN 
        likes_dislikes d ON d.user_id = ? AND d.comment_id = c.id AND d.value = false 
    LEFT JOIN (
        SELECT 
            comment_id, 
            COUNT(*) AS likes_count 
        FROM 
            likes_dislikes 
        WHERE 
            value = false 
        GROUP BY 
            comment_id
    ) ld ON ld.comment_id = c.id
    ORDER BY 
        l.id;
    `
	rows, err := repo.DB.Query(q, userId, userId)
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
			&comment.Likes,
			&comment.CreatedAt,
			&comment.Dislikes,
			&comment.Liked,
			&comment.Disliked,
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
