package comment

func (s *CommentService) EditComment(content string, commentId int, userId int) error {

	err := s.CommentRepo.UPDATE_comment(content, commentId, userId)
	if err != nil {
		return err
	}
	return nil
}
