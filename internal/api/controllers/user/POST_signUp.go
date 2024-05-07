package user

import (
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"forum/pkg/validation"
	"io"
	"log"
	"net/http"
	"net/mail"
)

func (m *UserController) POST_SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/signUp" {
		httpHelper.NotFoundError(w)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpHelper.InternalServerError(w)
		return
	}
	data := struct {
		Login           string `json:"login"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		httpHelper.BadRequestError(w)
		return
	}

	//////////////////////VALIDATION
	if data.ConfirmPassword != data.Password {
		httpHelper.BadRequestError(w)
		return
	}
	if data.Login == "" || data.Email == "" || data.Password == "" || data.ConfirmPassword == "" {
		httpHelper.BadRequestError(w)
		return
	}
	_, err = mail.ParseAddress(data.Email)
	if err != nil {
		httpHelper.BadRequestError(w)
		return
	}
	if !validation.PasswordIsValid(data.Password) {
		httpHelper.BadRequestError(w)
		return
	}
	//////////////////////VALIDATION
	err = m.UserService.RegisterUser(data.Login, data.Email, data.Password)
	if err != nil {
		switch err {
		case models.ErrConflict:
			httpHelper.ConflictError(w)
			return
		default:
			log.Println(err.Error())
			httpHelper.InternalServerError(w)
			return
		}
	}
	httpHelper.WriteJson(w, http.StatusOK, models.DefaultMessage{
		Message: "User registered :)",
	})
}
