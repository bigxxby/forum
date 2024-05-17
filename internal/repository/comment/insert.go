package comment

func (r *CommentRepo) INSERT_comment(userId, postId int, content string) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO comments (user_id , post_id , content ) VALUES ( ? , ? , ?)
	`
	res, err := r.DB.Exec(q, userId, postId, content)
	if err != nil {
		return -1, err
	}
	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}
	return int(commentId), nil
}
func (r *CommentRepo) INSERT_reply(userId, parentId int, content string) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO comments (user_id , content , parent_id) VALUES ( ? , ? , ?)
	`
	res, err := tx.Exec(q, userId, content, parentId)
	if err != nil {
		return -1, err
	}
	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}
	return int(commentId), nil
}
