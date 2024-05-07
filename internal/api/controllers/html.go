package controllers

import (
	"forum/internal/service"
	"forum/pkg/httpHelper"
	"net/http"
)

type HTMLController struct {
	HTMLService *service.HTMLService
}

func NewHTMLController(service *service.HTMLService) *HTMLController {
	return &HTMLController{
		HTMLService: service,
	}
}

func (c *HTMLController) GET_HTML_Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.ParseHTMLError(w, http.StatusMethodNotAllowed, "Method is not allowed, what are you trying to do? :}")
		return
	}
	if r.URL.Path != "/" {
		httpHelper.ParseHTMLError(w, 404, "Page not found... well maybe for now?")
		return
	}
	httpHelper.RenderHTMLPage(w, "web/ui/templates/index.html", nil)
}

func (c *HTMLController) GET_HTML_SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.ParseHTMLError(w, http.StatusMethodNotAllowed, "Method is not allowed, what are you trying to do? :}")
		return
	}
	if r.URL.Path != "/signUp" {
		httpHelper.ParseHTMLError(w, 404, "Page not found... well maybe for now?")
		return
	}
	httpHelper.RenderHTMLPage(w, "web/ui/templates/signUp.html", nil)
}

func (c *HTMLController) GET_HTML_SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.ParseHTMLError(w, http.StatusMethodNotAllowed, "Method is not allowed, what are you trying to do? :}")
		return
	}
	if r.URL.Path != "/signIn" {
		httpHelper.ParseHTMLError(w, 404, "Page not found... well maybe for now?")
		return
	}
	httpHelper.RenderHTMLPage(w, "web/ui/templates/signIn.html", nil)
}
