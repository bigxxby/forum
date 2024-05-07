package repository

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

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
func (r *UserRepository) CheckUserExists(email string) error {
	q := `SELECT COUNT(1) FROM users where email = ?`
	var exists int
	err := r.db.QueryRow(q, email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists != 0 {
		return errors.New("user elready exists")
	}
	return nil
}
func (r *UserRepository) CheckUserExistsByLogin(login string) error {
	q := `SELECT COUNT(1) FROM users where login = ?`
	var exists int
	err := r.db.QueryRow(q, login).Scan(&exists)
	if err != nil {
		return err
	}
	if exists != 0 {
		return errors.New("user elready exists")
	}
	return nil
}

func (r *UserRepository) CheckPassword(email, password string) error {
	q := `SELECT password FROM users WHERE email = ?`
	var hashPass string
	err := r.db.QueryRow(q, email).Scan(&hashPass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
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
