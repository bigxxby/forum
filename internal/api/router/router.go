package router

import (
	"database/sql"
	"forum/internal/api/controllers"
	"forum/internal/repository"
	"forum/internal/service"
	"log"
	"net/http"
)

type Router struct {
	UserController *controllers.UserController
	HTMLController *controllers.HTMLController
}

func NewRouter(connection *sql.DB) *Router {
	return &Router{
		UserController: controllers.NewUserController(service.NewUserService(repository.NewUserRepository(connection))),
		HTMLController: controllers.NewHTMLController(*service.NewHTMLService(repository.NewHTMLRepo(connection))),
	}
}

func Run() {
	connection, err := repository.CreateConnection()
	if err != nil {
		log.Println(err)
		return
	}
	err = repository.Drop(connection)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = repository.Migrate(connection)
	if err != nil {
		log.Println(err.Error())
		return
	}
	router := NewRouter(connection)

	http.HandleFunc("/", router.HTMLController.GET_HTML_Index)
	http.HandleFunc("/signUp", router.HTMLController.GET_HTML_SignUp)
	http.HandleFunc("/signIn", router.HTMLController.GET_HTML_SignIn)

	http.HandleFunc("/api/users/taken", router.UserController.GET_CheckIfLoginIsTaken)
	http.HandleFunc("/api/signUp", router.UserController.POST_SignUp)
	http.HandleFunc("/api/signIn", router.UserController.POST_SignIn)

	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	http.Handle(staticDir, staticFileServer)

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
