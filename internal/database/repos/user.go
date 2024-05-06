package repos

import (
	"database/sql"
	"errors"
	"forum/pkg/crypto"

	"golang.org/x/crypto/bcrypt"
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
	existsQ := `SELECT COUNT(1) FROM users where email = ?`
	var exists int
	err = r.db.QueryRow(existsQ, email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists != 0 {

		return errors.New("user elready exists")
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
func (r *UserRepo) AuthUser(email, password string) (string, error) {
	q := `SELECT password FROM users WHERE email = ?`
	var hashPass string
	err := r.db.QueryRow(q, email).Scan(&hashPass)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if err != nil {
		return "", err
	}
	uuid, err := crypto.CreateUUID()
	if err != nil {
		return "", err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	q = `UPDATE users SET uuid = ? WHERE email = ?`
	_, err = tx.Exec(q, uuid, email)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return uuid, err
}
