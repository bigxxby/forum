package user

import (
	"errors"
	"forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

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
func (r *UserRepository) SELECT_user(userId int) (*models.User, error) {
	q := `
	SELECT id ,  login , email FROM users WHERE id = ?
	`
	var user models.User
	err := r.db.QueryRow(q, userId).Scan(&user.ID, &user.Login, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
