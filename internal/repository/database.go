package repository

import (
	"database/sql"
	"os"
)

func Migrate(db *sql.DB) error {
	err := tables(db)
	if err != nil {
		return err
	}
	return nil
}
func Drop(db *sql.DB) error {
	migrationSQL, err := os.ReadFile("migrations/drop.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		return err
	}
	return nil
}
func tables(db *sql.DB) error {
	migrationSQL, err := os.ReadFile("migrations/users.sql")
	if err != nil {

		return err
	}
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		return err
	}
	return nil
}
