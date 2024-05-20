package user

import (
	"database/sql"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"log"
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
func (c *UserController) GET_MyProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/profile" {
		httpHelper.NotFoundError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, ok := userId.(int)
	if !ok {
		httpHelper.Unauthorised(w)
		return
	}

	user, err := c.UserService.GetMyProfile(userIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	if user == nil {
		httpHelper.NotFoundError(w)
		return
	}
	httpHelper.WriteJson(w, 200, user)
}
