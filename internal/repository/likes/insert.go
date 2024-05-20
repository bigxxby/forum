package likes

func (repo *LikesRepo) INSERT_like_post(userId, postId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO likes_dislikes (user_id , post_id  , value ) VALUES (? , ? , 1  )
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
func (repo *LikesRepo) INSERT_dislike_post(userId, postId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO likes_dislikes (user_id , post_id  , value ) VALUES (? , ? , 0  )
	
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
	INSERT INTO likes_dislikes (user_id , comment_id , value) VALUES (? , ?  , 1 )
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
func (repo *LikesRepo) INSERT_dislike_comment(userId, commentId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	INSERT INTO likes_dislikes (user_id , comment_id  , value ) VALUES (? , ? , 0 )
	
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
