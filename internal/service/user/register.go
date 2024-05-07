package user

import (
	"forum/internal/models"
	"forum/pkg/crypto"
)

func (s *UserService) RegisterUser(login, email, password string) error {
	err := s.UserRepository.CheckUserExists(email)
	if err != nil {
		return models.ErrConflict
	}
	err = s.UserRepository.CheckUserExistsByLogin(login)
	if err != nil {
		return models.ErrConflict
	}
	hash, err := crypto.GenerateHash(password)
	if err != nil {
		return err
	}
	err = s.UserRepository.InsertUser(login, hash, email)
	if err != nil {
		return err
	}
	return nil
}
