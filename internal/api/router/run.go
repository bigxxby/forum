package router

import (
	"database/sql"
	"forum/internal/api/middlewares"
	"forum/internal/repository"
	"log"
	"net/http"
)

// Run запускает HTTP сервер
func Run() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	connection, err := setupDatabase()
	if err != nil {
		log.Println(err)
		return
	}

	router := NewRouter(connection)
	router.Controller.UserController.UserService.CreateAdmin()

	mux := http.NewServeMux()

	// HTML маршруты
	registerHTMLRoutes(mux, router)

	// API маршруты
	registerAPIRoutes(mux, router)

	// Статические файлы
	registerStaticRoutes(mux)

	// Оборачиваем mux в логирующее middleware
	loggedMux := middlewares.LoggingMiddleware(mux)

	log.Println("Server started at http://localhost:8080/")
	err = http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

// setupDatabase инициализирует базу данных
func setupDatabase() (*sql.DB, error) {
	connection, err := repository.CreateConnection()
	if err != nil {
		return nil, err
	}
	err = repository.Drop(connection)
	if err != nil {
		return nil, err
	}
	err = repository.Migrate(connection)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// registerHTMLRoutes регистрирует маршруты для HTML страниц
func registerHTMLRoutes(mux *http.ServeMux, router *Router) {
	mux.HandleFunc("/", router.HTMLController.GET_HTML_Index)
	mux.HandleFunc("/signUp", router.HTMLController.GET_HTML_SignUp)
	mux.HandleFunc("/signIn", router.HTMLController.GET_HTML_SignIn)
	mux.HandleFunc("/posts", router.HTMLController.GET_HTML_Posts)
	mux.HandleFunc("/posts/", router.HTMLController.GET_HTML_Post)

	mux.HandleFunc("/posts/create", router.HTMLController.GET_HTML_Post)
}

// registerAPIRoutes регистрирует маршруты для API
func registerAPIRoutes(mux *http.ServeMux, router *Router) {
	// User API

	////GET
	mux.HandleFunc("/api/users/taken", router.Controller.UserController.GET_CheckIfLoginIsTaken)
	mux.HandleFunc("/api/profile", middlewares.AuthMiddleware(router.Controller.UserController.GET_MyProfile, router.Controller.UserController.UserService))

	////POST
	mux.HandleFunc("/api/signUp", router.Controller.UserController.POST_SignUp)
	mux.HandleFunc("/api/signIn", router.Controller.UserController.POST_SignIn)
	mux.HandleFunc("/api/logout", router.Controller.UserController.POST_Logout)

	// Post API
	////POST
	mux.HandleFunc("/api/posts", middlewares.AuthMiddleware(router.Controller.PostController.POST_PostPost, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/like", middlewares.AuthMiddleware(router.Controller.PostController.POST_Like, router.Controller.UserController.UserService))       // query ?valueId=number
	mux.HandleFunc("/api/posts/dislike", middlewares.AuthMiddleware(router.Controller.PostController.POST_DisLike, router.Controller.UserController.UserService)) // query ?valueId=number
	////GET
	mux.HandleFunc("/api/posts/one", middlewares.AuthMiddleware(router.Controller.PostController.GET_post, router.Controller.UserController.UserService)) // query ?valueId=number
	mux.HandleFunc("/api/posts/all", middlewares.AuthMiddleware(router.Controller.PostController.GET_posts, router.Controller.UserController.UserService))

	mux.HandleFunc("/api/posts/createdBy", middlewares.AuthMiddleware(router.Controller.PostController.GET_postsCreatedByUser, router.Controller.UserController.UserService)) // query ?valueId=number

	mux.HandleFunc("/api/posts/filter/category", middlewares.AuthMiddleware(router.Controller.PostController.GET_postsByCategory, router.Controller.UserController.UserService))

	mux.HandleFunc("/api/posts/liked", middlewares.AuthMiddleware(router.Controller.PostController.GET_likedPosts, router.Controller.UserController.UserService))

	// Comment API
	////POST
	mux.HandleFunc("/api/comments/post", middlewares.AuthMiddleware(router.Controller.CommentController.POST_Comment, router.Controller.UserController.UserService))    // query ?valueId=number
	mux.HandleFunc("/api/comments/like", middlewares.AuthMiddleware(router.Controller.CommentController.POST_Like, router.Controller.UserController.UserService))       // query ?valueId=number
	mux.HandleFunc("/api/comments/dislike", middlewares.AuthMiddleware(router.Controller.CommentController.POST_DisLike, router.Controller.UserController.UserService)) // query ?valueId=number
	mux.HandleFunc("/api/comments/edit", middlewares.AuthMiddleware(router.Controller.CommentController.PUT_EditComment, router.Controller.UserController.UserService)) // query ?valueId=number
	////DELETE
	mux.HandleFunc("/api/comments/delete", middlewares.AuthMiddleware(router.Controller.CommentController.DELETE_Comment, router.Controller.UserController.UserService)) // query ?valueId=number
	////GET
	mux.HandleFunc("/api/comments/liked", middlewares.AuthMiddleware(router.Controller.CommentController.GET_LikedComments, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments", middlewares.AuthMiddleware(router.Controller.CommentController.GET_Comments, router.Controller.UserController.UserService)) // query ?valueId=number

	// Category API
	////GET
	mux.HandleFunc("/api/categories", router.Controller.CategoryController.GET_categories)
}

// registerStaticRoutes регистрирует маршруты для статических файлов
func registerStaticRoutes(mux *http.ServeMux) {
	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	mux.Handle(staticDir, staticFileServer)
}
