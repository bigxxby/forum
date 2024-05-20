package likes

func (repo *LikesRepo) SELECT_alreadyLikedPost(userId, postId int) (bool, error) {
	q := `
	 SELECT EXISTS (
            SELECT 1 FROM likes_dislikes WHERE user_id = ? AND post_id = ? AND value = 1 
        )
	`
	var exists bool
	err := repo.DB.QueryRow(q, userId, postId).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}
func (repo *LikesRepo) SELECT_alreadyDisLikedPost(userId, postId int) (bool, error) {
	q := `
	 SELECT EXISTS (
            SELECT 1 FROM likes_dislikes WHERE user_id = ? AND post_id = ? AND value = 0
        )
	`
	var exists bool
	err := repo.DB.QueryRow(q, userId, postId).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}
func (repo *LikesRepo) SELECT_alreadyLikedComment(userId, commentId int) (bool, error) {
	q := `
	 SELECT EXISTS (
            SELECT 1 FROM likes_dislikes WHERE user_id = ? AND comment_id = ? AND value = 1
        )
	`
	var exists bool
	err := repo.DB.QueryRow(q, userId, commentId).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}
func (repo *LikesRepo) SELECT_alreadyDisLikedComment(userId, commentId int) (bool, error) {
	q := `
	 SELECT EXISTS (
            SELECT 1 FROM likes_dislikes WHERE user_id = ? AND comment_id = ? AND value = 0
        )
	`
	var exists bool
	err := repo.DB.QueryRow(q, userId, commentId).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return false, nil
}
