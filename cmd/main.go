package main

import (
	"forum/internal/api/router"

	_ "github.com/mattn/go-sqlite3"
)

// All standard Go packages are allowed.
// sqlite3
// bcrypt
// UUID
func main() {
	router.Run()
}
