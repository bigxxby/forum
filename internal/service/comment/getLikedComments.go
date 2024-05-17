package comment

import "forum/internal/models"

func (s *CommentService) GetAllLikedComments(userId int) ([]models.Comment, error) {
	comments, err := s.CommentRepo.SELECT_liked_comments(userId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
