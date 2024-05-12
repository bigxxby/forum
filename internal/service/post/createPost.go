package post

func (s *PostService) CreatePost(userId int, title, content string, categoryName string) (int, error) {
	categoryId, err := s.CategoryRepository.SELECT_categoryByName(categoryName)
	if err != nil {
		return -1, err
	}

	postId, err := s.PostRepository.INSERT_post(userId, title, content, categoryId)
	if err != nil {
		return -1, err
	}
	err = s.CategoryRepository.UPDATE_catCount(categoryId)
	if err != nil {
		return -1, err
	}
	return postId, nil
}
