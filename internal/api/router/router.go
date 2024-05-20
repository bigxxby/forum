package router

import (
	"database/sql"
	"forum/internal/api/controllers"
	"forum/internal/repository"
	"forum/internal/service"
)

type Router struct {
	Controller     *controllers.Controller
	HTMLController *controllers.HTMLController
}

func NewRouter(connection *sql.DB) *Router {
	repo := repository.NewRepository(connection)
	controller := controllers.NewController(service.NewService(repo))
	html := controllers.NewHTMLController(repo)
	return &Router{
		Controller:     controller,
		HTMLController: html,
	}
}
