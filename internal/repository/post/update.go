package post

func (repo *PostRepository) UPDATE_like(postId int, like bool) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := ``
	if like {
		q = `
	UPDATE posts SET likes = likes + 1 WHERE id = ?
	`
	} else {
		q = `
UPDATE posts 
SET likes = MAX(likes - 1, 0) WHERE id = ?
		`
	}
	_, err = tx.Exec(q, postId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
