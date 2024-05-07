package user

import (
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"net/http"
)

func (c *UserController) GET_CheckIfLoginIsTaken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/users/taken" {
		httpHelper.NotFoundError(w)
		return
	}
	login := r.URL.Query().Get("login")
	if login == "" {
		httpHelper.BadRequestError(w)
		return
	}

	err := c.UserService.CheckLoginAvailable(login)
	if err != nil {
		httpHelper.ConflictError(w)
		return
	}
	httpHelper.WriteJson(w, 200, models.DefaultMessage{
		Message: "Login is free :)",
	})
}
