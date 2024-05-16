package comment

func (repo *CommentRepo) UPDATE_like(commentId int, liked bool) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := ``
	if liked {
		q = ` 
		UPDATE comments SET likes = likes + 1  WHERE id = ? 
		`

	} else {
		q = ` 
		UPDATE comments SET  likes = MAX(likes - 1, 0 )  WHERE id =? 
		`
	}

	_, err = tx.Exec(q, commentId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return nil
	}
	return nil
}
