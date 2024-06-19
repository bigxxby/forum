package post

func (r *PostRepository) INSERT_post(userId int, title, content string, categoryIds []int) (int, error) {
	var postId int64
	res, err := r.DB.Exec("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", userId, title, content)

	if err != nil {
		return -1, err
	}
	postId, err = res.LastInsertId()
	if err != nil {
		return -1, err
	}

	for _, catId := range categoryIds {
		_, err = r.DB.Exec("INSERT INTO posts_categories (post_id, category_id) VALUES (?, ?)", postId, catId)
		if err != nil {
			return -1, err
		}
	}

	return int(postId), nil

}
