package user

import "forum/internal/models"

func (s *UserService) GetMyProfile(userId int) (*models.User, error) {
	user, err := s.UserRepository.SELECT_user(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
