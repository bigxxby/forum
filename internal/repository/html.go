package repository

import (
	"database/sql"
)

type HTMLRepository struct {
	DB *sql.DB
}

func NewHTMLRepo(db *sql.DB) *HTMLRepository {
	return &HTMLRepository{
		DB: db,
	}
}
