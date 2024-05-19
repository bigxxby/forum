package comment

import "database/sql"

func (r *CommentRepo) DELETE_Comment(userId int, commentId int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	q := `
	DELETE FROM comments WHERE user_id = ? AND id = ?
	`
	defer tx.Rollback()
	res, err := tx.Exec(q, userId, commentId)
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
