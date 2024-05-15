package comment

import "forum/internal/models"

func (s *CommentService) GetAllCommentsOfAPost(postId int) ([]models.Comment, error) {
	comments, err := s.CommentRepo.SELECT_Comments(postId)
	if err != nil {
		return nil, err
	}
	return comments, err

}
