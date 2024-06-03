package post

func (s *PostService) LikePost(userId, postId int) (string, error) {

	liked, err := s.LikesRepo.SELECT_alreadyLikedPost(userId, postId)
	if err != nil {
		return "", err
	}
	disliked, err := s.LikesRepo.SELECT_alreadyDisLikedPost(userId, postId)
	if err != nil {
		return "", err
	}
	if disliked {
		s.DisLikePost(userId, postId)
	}

	if !liked {
		err = s.PostRepository.UPDATE_like(postId, true)
		if err != nil {
			return "", err
		}
		err = s.LikesRepo.INSERT_like_post(userId, postId)
		if err != nil {
			return "", err
		}
	} else {
		err = s.PostRepository.UPDATE_like(postId, false)
		if err != nil {
			return "", err
		}
		err = s.LikesRepo.DELETE_unLike_post(userId, postId)
		if err != nil {
			return "", err
		}
	}
	if liked {
		return "UnLiked :(", nil
	} else {
		return "Liked :)", nil

	}
}
func (s *PostService) DisLikePost(userId, postId int) (string, error) {

	disliked, err := s.LikesRepo.SELECT_alreadyDisLikedPost(userId, postId)
	if err != nil {
		return "", err
	}
	liked, err := s.LikesRepo.SELECT_alreadyLikedPost(userId, postId)
	if err != nil {
		return "", err

	}
	if liked {
		s.LikePost(userId, postId)
	}
	if !disliked {
		err = s.PostRepository.UPDATE_dislike(postId, true)
		if err != nil {
			return "", err
		}
		err = s.LikesRepo.INSERT_dislike_post(userId, postId)
		if err != nil {
			return "", err
		}
	} else {
		err = s.PostRepository.UPDATE_dislike(postId, false)
		if err != nil {
			return "", err
		}
		err = s.LikesRepo.DELETE_unDisLike_post(userId, postId)
		if err != nil {
			return "", err
		}
	}
	if disliked {
		return "UnDisliked :)", nil
	} else {
		return "Disliked :(", nil

	}
}
