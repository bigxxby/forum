package user

import (
	"forum/internal/models"
	"forum/pkg/crypto"
)

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
