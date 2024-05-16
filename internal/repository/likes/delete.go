package likes

func (repo *LikesRepo) DELETE_unLike_post(userId, postId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	DELETE FROM likes WHERE user_id = ? AND post_id = ? 
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
func (repo *LikesRepo) DELETE_unLike_comment(userId, commentId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `
	DELETE FROM likes WHERE user_id = ? AND comment_id = ? 
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
