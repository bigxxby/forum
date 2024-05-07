package user

func (r *UserRepository) InsertUser(login, hash, email string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `INSERT INTO users (login , email , password) VALUES ( ? , ? , ?)`
	_, err = tx.Exec(q, login, email, hash)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) InsertAdmin(hash string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `INSERT INTO users (login , email , password) VALUES ( ? , ? , ?)`
	_, err = tx.Exec(q, "bigxxby", "bigxxby@yandex.ru", hash)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
