package comment

import (
	"database/sql"
)

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

	res, err := tx.Exec(q, commentId)
	if err != nil {
		return err
	}

	last, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if last == 0 {
		return sql.ErrNoRows
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (repo *CommentRepo) UPDATE_dislike(commentId int, liked bool) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := ``
	if liked {
		q = ` 
		UPDATE comments SET dislikes = dislikes + 1  WHERE id = ? 
		`

	} else {
		q = ` 
		UPDATE comments SET  dislikes = MAX(dislikes - 1, 0 )  WHERE id =? 
		`
	}

	res, err := tx.Exec(q, commentId)
	if err != nil {
		return err
	}

	last, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if last == 0 {
		return sql.ErrNoRows
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
