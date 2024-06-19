package category

import (
	"strings"
)

func (repo *CategoryRepository) SELECT_categories() (map[string]int, error) {
	q := `SELECT name , posts_count  FROM categories`

	rows, err := repo.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := make(map[string]int)
	for rows.Next() {
		var category string
		var count int
		err := rows.Scan(&category, &count)
		if err != nil {
			return nil, err
		}
		categories[category] = count
	}
	return categories, nil
}

func (repo *CategoryRepository) SELECT_categoriesByName(categoryName []string) ([]int, error) {
	for i, cat := range categoryName {
		categoryName[i] = "'" + cat + "'"
	}
	q := `SELECT id FROM categories WHERE name IN (` + strings.Join(categoryName, ",") + `)`
	rows, err := repo.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var catIds []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		catIds = append(catIds, id)
	}
	return catIds, nil

}
