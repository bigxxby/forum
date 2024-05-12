package category

func (repo *CategoryRepository) UPDATE_catCount(catId int) error {
	// Начать транзакцию
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
		UPDATE categories 
		SET posts_count = posts_count + 1
		WHERE id = ?
	`
	_, err = tx.Exec(q, catId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
