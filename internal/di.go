package internal

import (
	"forum/internal/api/controllers"
	"forum/internal/database"
)

type DIstruct struct {
	Database *database.Database
	Manager  *controllers.Manager
}

func NewDiStruct() (*DIstruct, error) {
	connection, err := database.CreateConnection()
	if err != nil {
		return nil, err
	}
	newDatabase := database.Database{
		DB: connection,
	}

	manager, err := controllers.NewManager(newDatabase)
	if err != nil {
		return nil, err
	}
	di := DIstruct{
		Database: &newDatabase,
		Manager:  manager,
	}
	return &di, nil
}
