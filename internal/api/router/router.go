package router

import (
	"forum/internal"
	"log"
	"net/http"
)

func Run() {
	diStruct, err := internal.NewDiStruct()
	log.Println(diStruct.Manager)
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.HandleFunc("/", diStruct.Manager.GET_HelloWorld)

	http.ListenAndServe(":8080", nil)

}
