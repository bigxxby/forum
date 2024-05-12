package post

func (r *PostRepository) INSERT_post(userId int, title, content string) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	q := `INSERT INTO posts (user_id , title , content ) VALUES( ?,?, ?)`
	res, err := tx.Exec(q, userId, title, content)
	if err != nil {
		return -1, err
	}

	postId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, nil
	}
	return int(postId), nil
}
