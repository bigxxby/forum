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
	controller := controllers.NewController(service.NewService(repository.NewRepository(connection)))
	html := controllers.NewHTMLController(connection)
	return &Router{
		Controller:     controller,
		HTMLController: html,
	}
}
