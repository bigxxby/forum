package controllers

import (
	"forum/internal/repository"
	"forum/pkg/httpHelper"
	"net/http"
)

type HTMLController struct {
	Repo *repository.Repository
}

func NewHTMLController(r *repository.Repository) *HTMLController {
	return &HTMLController{
		Repo: r,
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
func (c *HTMLController) GET_HTML_Posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.ParseHTMLError(w, http.StatusMethodNotAllowed, "Method is not allowed, what are you trying to do? :}")
		return
	}
	httpHelper.RenderHTMLPage(w, "web/ui/templates/posts.html", nil)
}
func (c *HTMLController) GET_HTML_Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.ParseHTMLError(w, http.StatusMethodNotAllowed, "Method is not allowed, what are you trying to do? :}")
		return
	}
	uuid, err := r.Cookie("uuid")
	if err != nil {
		httpHelper.ParseHTMLError(w, http.StatusUnauthorized, "Unathorised, LOGIN FIRST!!!")
		return
	}

	user, err := c.Repo.UserRepository.GetUserByUUID(uuid.Value)
	if err != nil {
		if err != nil {
			httpHelper.ParseHTMLError(w, http.StatusUnauthorized, "Unathorised, LOGIN FIRST!!!")
			return
		}
	}
	if user == nil {
		httpHelper.ParseHTMLError(w, http.StatusUnauthorized, "Unathorised, LOGIN FIRST!!!")
		return
	}
	httpHelper.RenderHTMLPage(w, "web/ui/templates/createPost.html", nil)
}
