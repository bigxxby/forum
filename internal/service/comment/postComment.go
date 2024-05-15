package comment

import (
	"errors"
	"forum/pkg/validation"
)

func (s *CommentService) PostComment(userId, postId int, content string) (int, error) {
	if !validation.IsValidComment(content) {
		return -1, errors.New("comment is not valid")
	}
	commentId, err := s.CommentRepo.INSERT_comment(userId, postId, content)
	if err != nil {
		return -1, err
	}
	return commentId, nil
}
