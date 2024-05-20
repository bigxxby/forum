package post

import "database/sql"

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
	res, err := tx.Exec(q, postId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (repo *PostRepository) UPDATE_dislike(postId int, like bool) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := ``
	if like {
		q = `
	UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?
	`
	} else {
		q = `
UPDATE posts 
SET dislikes = MAX(dislikes - 1, 0) WHERE id = ?
		`
	}
	res, err := tx.Exec(q, postId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
