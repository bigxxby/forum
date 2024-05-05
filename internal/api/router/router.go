package router

import (
	"forum/internal"
	"forum/internal/database"
	"log"
	"net/http"
)

func Run() {
	diStruct, err := internal.NewDiStruct()
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.Drop(diStruct.Database.DB)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = database.Migrate(diStruct.Database.DB)
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.HandleFunc("/api/signUp", diStruct.Manager.POST_SignUp)

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
