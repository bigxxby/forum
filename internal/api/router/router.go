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

	http.HandleFunc("/", diStruct.Manager.GET_HTML_Index)
	http.HandleFunc("/signUp", diStruct.Manager.GET_HTML_SignUp)
	http.HandleFunc("/signIn", diStruct.Manager.GET_HTML_SignIn)

	http.HandleFunc("/api/signUp", diStruct.Manager.POST_SignUp)
	http.HandleFunc("/api/signIn", diStruct.Manager.POST_SignIn)

	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	http.Handle(staticDir, staticFileServer)

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
