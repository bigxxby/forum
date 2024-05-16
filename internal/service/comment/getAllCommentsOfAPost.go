package comment

import "forum/internal/models"

func (s *CommentService) GetAllCommentsOfAPost(postId int, userId int) ([]models.Comment, error) {
	comments, err := s.CommentRepo.SELECT_Comments(postId, userId)
	if err != nil {
		return nil, err
	}
	return comments, err

}
