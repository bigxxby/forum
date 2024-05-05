package repos

import (
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) CreateUser(uuid, login, hash, email string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := `INSERT INTO users (uuid , login , email , password) VALUES ( ? , ? , ? , ?)`
	_, err = tx.Exec(q, uuid, login, email, hash)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
