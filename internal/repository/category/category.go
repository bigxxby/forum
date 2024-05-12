package category

import "database/sql"

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(connection *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: connection,
	}
}
