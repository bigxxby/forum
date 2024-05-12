package category

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

func (repo *CategoryRepository) SELECT_categoryByName(categoryName string) (int, error) {
	q := `SELECT id FROM categories WHERE name = ?`

	var id int
	err := repo.DB.QueryRow(q, categoryName).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil

}
