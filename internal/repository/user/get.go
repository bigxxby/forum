package user

import "forum/internal/models"

func (r *UserRepository) GetUserByUUID(uuid string) (*models.User, error) {
	var user models.User
	q := `SELECT  id , login , email  FROM users where uuid = ?`
	err := r.db.QueryRow(q, uuid).Scan(&user.ID, &user.Login, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
