package comment

func (s *CommentService) DeleteComment(userId int, commentId int) error {
	err := s.CommentRepo.DELETE_Comment(userId, commentId)
	if err != nil {
		return err
	}
	return nil
}
