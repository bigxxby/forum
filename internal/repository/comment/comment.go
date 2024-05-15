package comment

import "database/sql"

type CommentRepo struct {
	DB *sql.DB
}

func NewCommentRepo(conn *sql.DB) *CommentRepo {
	return &CommentRepo{
		DB: conn,
	}
}
