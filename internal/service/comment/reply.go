package comment

import (
	"errors"
	"forum/pkg/validation"
)

func (s *CommentService) ReplyToComment(parentId int, userId int, content string) (int, error) {
	valid := validation.IsValidComment(content)
	if !valid {
		return -1, errors.New("comment is not valid")
	}

	commentId, err := s.CommentRepo.INSERT_reply(userId, parentId, content)
	if err != nil {
		return -1, err
	}
	return commentId, nil
}
