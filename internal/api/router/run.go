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
	registerStaticRoutes()

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
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

	// Post API
	////POST
	mux.HandleFunc("/api/posts", middlewares.AuthMiddleware(router.Controller.PostController.POST_PostPost, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/{id}/like", middlewares.AuthMiddleware(router.Controller.PostController.POST_Like, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/{id}/dislike", middlewares.AuthMiddleware(router.Controller.PostController.POST_DisLike, router.Controller.UserController.UserService))
	////GET
	mux.HandleFunc("/api/posts/{id}", middlewares.AuthMiddleware(router.Controller.PostController.GET_post, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/all", middlewares.AuthMiddleware(router.Controller.PostController.GET_posts, router.Controller.UserController.UserService))

	mux.HandleFunc("/api/posts/{userId}/createdBy", middlewares.AuthMiddleware(router.Controller.PostController.GET_postsCreatedByUser, router.Controller.UserController.UserService))

	mux.HandleFunc("/api/posts/filter/category", middlewares.AuthMiddleware(router.Controller.PostController.GET_postsByCategory, router.Controller.UserController.UserService))

	mux.HandleFunc("/api/posts/liked", middlewares.AuthMiddleware(router.Controller.PostController.GET_likedPosts, router.Controller.UserController.UserService))

	// Comment API
	////POST
	mux.HandleFunc("/api/comments/{id}/post", middlewares.AuthMiddleware(router.Controller.CommentController.POST_Comment, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{id}/like", middlewares.AuthMiddleware(router.Controller.CommentController.POST_Like, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{id}/dislike", middlewares.AuthMiddleware(router.Controller.CommentController.POST_DisLike, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{id}/edit", middlewares.AuthMiddleware(router.Controller.CommentController.PUT_EditComment, router.Controller.UserController.UserService))
	////DELETE
	mux.HandleFunc("/api/comments/{id}/delete", middlewares.AuthMiddleware(router.Controller.CommentController.DELETE_Comment, router.Controller.UserController.UserService))
	////GET
	mux.HandleFunc("/api/comments/liked", middlewares.AuthMiddleware(router.Controller.CommentController.GET_LikedComments, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{postId}", middlewares.AuthMiddleware(router.Controller.CommentController.GET_Comments, router.Controller.UserController.UserService))

	// Category API
	////GET
	mux.HandleFunc("/api/categories", router.Controller.CategoryController.GET_categories)
}

// registerStaticRoutes регистрирует маршруты для статических файлов
func registerStaticRoutes() {
	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	http.Handle(staticDir, staticFileServer)
}
