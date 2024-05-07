package user

import "forum/pkg/crypto"

func (s *UserService) CreateAdmin() error {
	hash, err := crypto.GenerateHash("Aa123123")
	if err != nil {
		return err
	}
	err = s.UserRepository.InsertAdmin(hash)
	if err != nil {
		return err
	}
	return nil
}
