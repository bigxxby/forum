package user

import "forum/internal/models"

func (s *UserService) CheckLoginAvailable(login string) error {
	err := s.UserRepository.CheckUserExistsByLogin(login)
	if err != nil {
		return models.ErrConflict
	}
	return nil
}
