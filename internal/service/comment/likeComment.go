package comment

import (
	"errors"
)

func (s *CommentService) LikeComment(userId int, commentId int) error {
	liked, err := s.LikeRepo.SELECT_alreadyLikedComment(userId, commentId)
	if err != nil {
		return err
	}
	if liked {
		return errors.New("comment already liked")
	}

	err = s.CommentRepo.UPDATE_like(commentId, true)
	if err != nil {
		return err
	}
	err = s.LikeRepo.INSERT_like_comment(userId, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentService) UnLikeComment(userId int, commentId int) error { //
	liked, err := s.LikeRepo.SELECT_alreadyLikedComment(userId, commentId)
	if err != nil {
		return err
	}
	if !liked {
		return errors.New("comment not liked")
	}
	err = s.CommentRepo.UPDATE_like(commentId, false)
	if err != nil {
		return err
	}
	err = s.LikeRepo.DELETE_unLike_comment(userId, commentId)
	if err != nil {
		return err
	}
	return nil
}
