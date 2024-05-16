package likes

import (
	"database/sql"
)

type LikesRepo struct {
	DB *sql.DB
}

func NewLikesRepo(connection *sql.DB) *LikesRepo {
	return &LikesRepo{
		DB: connection,
	}
}
