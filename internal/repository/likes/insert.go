package likes

func (repo *LikesRepo) INSERT_like_post(userId, postId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO likes (user_id , post_id ) VALUES (? , ? )
	
	`
	_, err = tx.Exec(q, userId, postId)

	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (repo *LikesRepo) INSERT_like_comment(userId, commentId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO likes (user_id , comment_id ) VALUES (? , ? )
	
	`
	_, err = tx.Exec(q, userId, commentId)

	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
