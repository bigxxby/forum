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
	// err := repository.
	mux := http.NewServeMux()
	router := NewRouter(connection)

	router.Controller.UserController.UserService.CreateAdmin()

	mux.HandleFunc("/", router.HTMLController.GET_HTML_Index)
	mux.HandleFunc("/signUp", router.HTMLController.GET_HTML_SignUp)
	mux.HandleFunc("/signIn", router.HTMLController.GET_HTML_SignIn)

	//TODO:
	// http.HandleFunc("/posts/{ID}", router.HTMLController.GET_HTML_SignIn)

	mux.HandleFunc("/api/users/taken", router.Controller.UserController.GET_CheckIfLoginIsTaken)
	mux.HandleFunc("/api/signUp", router.Controller.UserController.POST_SignUp)
	mux.HandleFunc("/api/signIn", router.Controller.UserController.POST_SignIn)

	mux.HandleFunc("/api/posts", middlewares.AuthMiddleware(router.Controller.PostController.POST_PostPost, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/{id}", middlewares.AuthMiddleware(router.Controller.PostController.GET_post, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/all", middlewares.AuthMiddleware(router.Controller.PostController.GET_posts, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/{id}/like", middlewares.AuthMiddleware(router.Controller.PostController.POST_Like, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/{id}/unlike", middlewares.AuthMiddleware(router.Controller.PostController.POST_Unlike, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/posts/liked", middlewares.AuthMiddleware(router.Controller.PostController.GET_likedPosts, router.Controller.UserController.UserService))

	mux.HandleFunc("/api/categories", router.Controller.CategoryController.GET_categories)

	mux.HandleFunc("/api/comments/{id}/post", middlewares.AuthMiddleware(router.Controller.CommentController.POST_Comment, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{id}/like", middlewares.AuthMiddleware(router.Controller.CommentController.POST_Like, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/liked", middlewares.AuthMiddleware(router.Controller.CommentController.GET_LikedComments, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{id}/unlike", middlewares.AuthMiddleware(router.Controller.CommentController.POST_UnLike, router.Controller.UserController.UserService))
	mux.HandleFunc("/api/comments/{id}", middlewares.AuthMiddleware(router.Controller.CommentController.GET_Comments, router.Controller.UserController.UserService))

	// mux.HandleFunc("/api/likes/posts", middlewares.AuthMiddleware(router.Controller.CommentController.GET_Comments, router.Controller.UserController.UserService))
	// mux.HandleFunc("/api/likes/comments", middlewares.AuthMiddleware(router.Controller.CommentController.GET_Comments, router.Controller.UserController.UserService))

	staticDir := "/static/"
	staticFileServer := http.StripPrefix(staticDir, http.FileServer(http.Dir("web/ui/static")))
	http.Handle(staticDir, staticFileServer)

	log.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
