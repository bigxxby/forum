package post

import (
	"database/sql"
	"forum/internal/models"
)

func (repo *PostRepository) SELECT_post(postId, userId int) (*models.Post, bool, bool, error) {
	q := `
    SELECT 
        p.id, 
        p.user_id, 
        p.title, 
        p.content, 
        p.created_at, 
        u.login,
        p.likes, 
        c.name,
        l1.value AS like_value,
        l2.value AS dislike_value,
        CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
        CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked,
		p.dislikes
		FROM 
        posts p
    LEFT JOIN 
        users u ON p.user_id = u.id
    LEFT JOIN 
        categories c ON p.category_id = c.id
    LEFT JOIN 
        likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
    LEFT JOIN 
        likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
    WHERE 
        p.id = ?;
    `
	var post models.Post
	var likeValue sql.NullBool
	var dislikeValue sql.NullBool
	var liked bool
	var disliked bool

	err := repo.DB.QueryRow(q, userId, userId, postId).Scan(
		&post.ID,
		&post.UserId,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.CreatedBy,
		&post.Likes,
		&post.Category,
		&likeValue,
		&dislikeValue,
		&liked,
		&disliked,
		&post.Dislikes,
	)
	if err != nil {
		return nil, false, false, err
	}

	return &post, liked, disliked, nil
}

func (repo *PostRepository) SELECT_postsCreatedAt(userId int) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        ORDER BY 
            p.created_at DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
func (repo *PostRepository) SELECT_postsMostLiked(userId int) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        ORDER BY 
            p.likes DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
func (repo *PostRepository) SELECT_postsMostDisliked(userId int) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        ORDER BY 
            p.dislikes DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
func (repo *PostRepository) SELECT_createdByThisUser(userId int, userIdCreatedBy int) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
		WHERE p.user_id = ?
        ORDER BY 
            p.created_at DESC
		;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId, userIdCreatedBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *PostRepository) SELECT_liked_posts(userId int) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.value IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.value IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
            l1.post_id IS NOT NULL
        ORDER BY 
            l1.id;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
func (repo *PostRepository) SELECT_postsByCategory(userId int, categoryName string) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.value IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.value IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
			c.name = ?
		ORDER BY 
            p.created_at;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
func (repo *PostRepository) SELECT_postsByMostLiked(userId int, categoryName string) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.value IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.value IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
			c.name = ?
		ORDER BY 
            p.likes DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
func (repo *PostRepository) SELECT_postsByMostDisliked(userId int, categoryName string) ([]models.Post, error) {
	q := `
        SELECT 
            p.id, 
            p.user_id, 
            p.title, 
            p.content, 
            p.created_at, 
            u.login, 
            p.likes, 
            c.name,
            p.dislikes,
            CASE WHEN l1.value IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.value IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            categories c ON p.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
			c.name = ?
		ORDER BY 
            p.dislikes DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var liked bool
		var disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Category,
			&post.Dislikes,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
