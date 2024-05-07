package post

import "database/sql"

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(connection *sql.DB) *PostRepository {
	return &PostRepository{
		DB: connection,
	}
}
