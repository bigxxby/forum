package post

import (
	"database/sql"
	"forum/internal/models"
	"strings"
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
        p.dislikes,
        COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
        CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
        CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
    FROM 
        posts p
    LEFT JOIN 
        users u ON p.user_id = u.id
    LEFT JOIN 
        posts_categories pc ON p.id = pc.post_id
    LEFT JOIN 
        categories c ON pc.category_id = c.id
    LEFT JOIN 
        likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
    LEFT JOIN 
        likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
    WHERE 
        p.id = ?
    GROUP BY 
        p.id, 
        p.user_id, 
        p.title, 
        p.content, 
        p.created_at, 
        u.login, 
        p.likes, 
        p.dislikes, 
        l1.post_id, 
        l2.post_id;
    `
	var post models.Post
	var categories string
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
		&post.Dislikes,
		&categories,
		&liked,
		&disliked,
	)
	if err != nil {
		return nil, false, false, err
	}

	// Split the concatenated categories into a slice
	post.Categories = strings.Split(categories, ",")

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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var createdBy sql.NullString
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&createdBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.CreatedBy = createdBy.String // Handle NULL-to-string conversion
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil // Handle empty category case
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
            p.user_id = ?
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
        ORDER BY 
            p.created_at DESC;
    `
	var posts []models.Post

	rows, err := repo.DB.Query(q, userId, userId, userIdCreatedBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
            l1.post_id IS NOT NULL
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c2.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            posts_categories pc2 ON p.id = pc2.post_id
        LEFT JOIN 
            categories c2 ON pc2.category_id = c2.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
            c.name = ?
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
            c.name = ?
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
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
            p.dislikes,
            COALESCE(GROUP_CONCAT(c.name, ','), '') AS categories,
            CASE WHEN l1.post_id IS NOT NULL THEN true ELSE false END AS liked,
            CASE WHEN l2.post_id IS NOT NULL THEN true ELSE false END AS disliked
        FROM 
            posts p
        LEFT JOIN 
            users u ON p.user_id = u.id
        LEFT JOIN 
            posts_categories pc ON p.id = pc.post_id
        LEFT JOIN 
            categories c ON pc.category_id = c.id
        LEFT JOIN 
            likes_dislikes l1 ON p.id = l1.post_id AND l1.user_id = ? AND l1.value = true
        LEFT JOIN 
            likes_dislikes l2 ON p.id = l2.post_id AND l2.user_id = ? AND l2.value = false
        WHERE 
            c.name = ?
        GROUP BY 
            p.id, p.user_id, p.title, p.content, p.created_at, u.login, p.likes, p.dislikes
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
		var categories string
		var liked, disliked bool

		err := rows.Scan(
			&post.ID,
			&post.UserId,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.CreatedBy,
			&post.Likes,
			&post.Dislikes,
			&categories,
			&liked,
			&disliked,
		)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ",")
		if categories == "" {
			post.Categories = nil
		}
		post.Liked = liked
		post.Disliked = disliked
		posts = append(posts, post)
	}
	return posts, nil
}
