package user

func (r *UserRepository) UpdateUUID(uuid, email string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	q := `UPDATE users SET uuid = ? WHERE email = ?`
	_, err = tx.Exec(q, uuid, email)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}
