package post

func (s *PostService) CreatePost(userId int, title, content string) (int, error) {
	postId, err := s.PostRepository.CreatePost(userId, title, content)
	if err != nil {
		return -1, err
	}
	return postId, nil
}
