package router

import (
	"database/sql"
	"forum/internal/api/controllers"
	postController "forum/internal/api/controllers/post"
	userController "forum/internal/api/controllers/user"
	"forum/internal/api/middlewares"
	"forum/internal/repository"
	"forum/internal/repository/post"
	userRepository "forum/internal/repository/user"
	"forum/internal/service"
	postService "forum/internal/service/post"
	userService "forum/internal/service/user"
	"log"
	"net/http"
)

type Router struct {
	UserController *userController.UserController
	HTMLController *controllers.HTMLController
	PostController *postController.PostController
}

func NewRouter(connection *sql.DB) *Router {
	userRepo := userRepository.NewUserRepository(connection)
	userServ := userService.NewUserService(userRepo)
	userCtrl := userController.NewUserController(userServ)

	htmlRepo := repository.NewHTMLRepo(connection)
	htmlServ := service.NewHTMLService(htmlRepo)
	htmlCtrl := controllers.NewHTMLController(htmlServ)

	postRepo := post.NewPostRepository(connection)
	postServ := postService.NewPostService(postRepo)
	postCtrl := postController.NewPostController(postServ, userServ)

	return &Router{
		UserController: userCtrl,
		HTMLController: htmlCtrl,
		PostController: postCtrl,
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
	err = router.UserController.UserService.CreateAdmin()
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.HandleFunc("/", router.HTMLController.GET_HTML_Index)
	http.HandleFunc("/signUp", router.HTMLController.GET_HTML_SignUp)
	http.HandleFunc("/signIn", router.HTMLController.GET_HTML_SignIn)

	http.HandleFunc("/api/users/taken", router.UserController.GET_CheckIfLoginIsTaken)
	http.HandleFunc("/api/signUp", router.UserController.POST_SignUp)
	http.HandleFunc("/api/signIn", router.UserController.POST_SignIn)

	http.HandleFunc("/api/posts", middlewares.AuthMiddleware(router.PostController.POST_PostPost, router.UserController.UserService))
	http.HandleFunc("/api/posts/", router.PostController.GET_post)

	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	http.Handle(staticDir, staticFileServer)

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
