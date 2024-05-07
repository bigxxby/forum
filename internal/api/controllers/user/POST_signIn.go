package user

import (
	"database/sql"
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"io"
	"log"
	"net/http"
)

func (m *UserController) POST_SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/signIn" {
		httpHelper.NotFoundError(w)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpHelper.InternalServerError(w)
		return
	}
	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		httpHelper.BadRequestError(w)
		return
	}
	//////////////////////VALIDATION
	if data.Email == "" || data.Password == "" {
		httpHelper.BadRequestError(w)
		return
	}
	//////////////////////VALIDATION
	uuid, err := m.UserService.AuthUser(data.Email, data.Password)
	if err != nil {
		if err == sql.ErrNoRows || err == models.ErrInvalidCredentials {
			httpHelper.Unauthorised(w)
			return
		} else {
			log.Println(err.Error())
			httpHelper.InternalServerError(w)
			return
		}
	}
	COCK := &http.Cookie{
		Name:     "uuid",
		Value:    uuid,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	http.SetCookie(w, COCK)

	httpHelper.WriteJson(w, http.StatusOK, models.AuthenticationMessage{
		Message: "Congratz, authorised :)",
		UUID:    uuid,
	})
}
