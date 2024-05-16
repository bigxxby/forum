package post

import (
	"errors"
)

func (s *PostService) LikePost(userId, postId int) error {

	liked, err := s.LikesRepo.SELECT_alreadyLikedPost(userId, postId)
	if err != nil {
		return err
	}

	if liked {

		return errors.New("post already liked")
	}

	err = s.PostRepository.UPDATE_like(postId, true)
	if err != nil {
		return err
	}
	err = s.LikesRepo.INSERT_like_post(userId, postId)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostService) UnLikePost(userId, postId int) error {

	liked, err := s.LikesRepo.SELECT_alreadyLikedPost(userId, postId)
	if err != nil {
		return err
	}
	if !liked {
		return errors.New("post not liked")
	}
	err = s.PostRepository.UPDATE_like(postId, false)
	if err != nil {
		return err
	}
	err = s.LikesRepo.DELETE_unLike_post(userId, postId)
	if err != nil {
		return err
	}
	return nil
}
