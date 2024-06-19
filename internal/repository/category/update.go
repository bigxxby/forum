package category

import (
	"fmt"
	"strconv"
	"strings"
)

func (repo *CategoryRepository) UPDATE_catCount(catId []int) error {
	// Convert []int to []string for SQL query
	var strIds []string
	for _, id := range catId {
		strIds = append(strIds, strconv.Itoa(id))
	}

	// Prepare the SQL query
	q := `UPDATE categories SET posts_count = posts_count + 1 WHERE id IN (` + strings.Join(strIds, ",") + `)`

	// Execute the query
	_, err := repo.DB.Exec(q)
	if err != nil {
		return fmt.Errorf("error updating category count: %w", err)
	}

	return nil
}
