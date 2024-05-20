package comment

func (s *CommentService) LikeOrUnlikeComment(userId int, commentId int) (string, error) {
	liked, err := s.LikeRepo.SELECT_alreadyLikedComment(userId, commentId)
	if err != nil {
		return "", err
	}

	if !liked {
		err = s.CommentRepo.UPDATE_like(commentId, true)
		if err != nil {
			return "", err
		}
		err = s.LikeRepo.INSERT_like_comment(userId, commentId)
		if err != nil {
			return "", err
		}
	} else {
		err = s.CommentRepo.UPDATE_like(commentId, false)
		if err != nil {
			return "", err
		}
		err = s.LikeRepo.DELETE_unLike_comment(userId, commentId)
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
func (s *CommentService) DisLikeOrUnDislikeComment(userId int, commentId int) (string, error) {
	disliked, err := s.LikeRepo.SELECT_alreadyDisLikedComment(userId, commentId)
	if err != nil {
		return "", err
	}

	if !disliked {
		err = s.CommentRepo.UPDATE_dislike(commentId, true)
		if err != nil {
			return "", err
		}
		err = s.LikeRepo.INSERT_dislike_comment(userId, commentId)
		if err != nil {
			return "", err
		}
	} else {
		err = s.CommentRepo.UPDATE_dislike(commentId, false)
		if err != nil {
			return "", err
		}
		err = s.LikeRepo.DELETE_unDisLike_comment(userId, commentId)
		if err != nil {
			return "", err
		}

	}
	if disliked {
		return "UnDisLiked :)", nil
	} else {
		return "DisLiked :(", nil

	}

}
