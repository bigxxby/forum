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
func (repo *CommentRepo) UPDATE_comment(content string, commentId int, userId int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	q := `
	UPDATE comments SET content = ? , edited = 1 WHERE id = ? AND user_id = ?
	`

	res, err := tx.Exec(q, content, commentId, userId)
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
	defer tx.Rollback()
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
