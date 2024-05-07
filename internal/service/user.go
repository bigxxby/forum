package service

import (
	"forum/internal/models"
	"forum/internal/repository"
	"forum/pkg/crypto"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (s *UserService) AuthUser(email, password string) (string, error) {
	//checks if given creditnails is valid
	err := s.UserRepository.CheckPassword(email, password)
	if err != nil {
		return "", models.ErrInvalidCredentials
	}
	// if all correct gives UUID token
	uuid, err := crypto.GenerateUUID()
	if err != nil {
		return "", models.ErrInternalServer
	}
	// and writes it to db
	err = s.UserRepository.UpdateUUID(uuid, email)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
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
func (s *UserService) CheckLoginAvailable(login string) error {
	err := s.UserRepository.CheckUserExistsByLogin(login)
	if err != nil {
		return models.ErrConflict
	}
	return nil
}
