package router

import (
	"forum/internal/api/middlewares"
	"forum/internal/repository"
	"log"
	"net/http"
)

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

	//TODO:
	// http.HandleFunc("/posts/{ID}", router.HTMLController.GET_HTML_SignIn)

	http.HandleFunc("/api/users/taken", router.Controller.UserController.GET_CheckIfLoginIsTaken)
	http.HandleFunc("/api/signUp", router.Controller.UserController.POST_SignUp)
	http.HandleFunc("/api/signIn", router.Controller.UserController.POST_SignIn)

	http.HandleFunc("/api/posts", middlewares.AuthMiddleware(router.Controller.PostController.POST_PostPost, router.Controller.UserController.UserService))
	http.HandleFunc("/api/posts/", router.Controller.PostController.GET_post)
	http.HandleFunc("/api/posts/all", router.Controller.PostController.GET_posts)

	http.HandleFunc("/api/categories", router.Controller.CategoryController.GET_categories)

	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	http.Handle(staticDir, staticFileServer)

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
