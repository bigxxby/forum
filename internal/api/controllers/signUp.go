package controllers

import (
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/crypto"
	"forum/pkg/httpHelper"
	"forum/pkg/validation"
	"io"
	"log"
	"net/http"
	"net/mail"
)

func (m *Manager) POST_SignUp(w http.ResponseWriter, r *http.Request) {
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
		Login           string `json:"login" required:"true"`
		Email           string `json:"email" required:"true"`
		Password        string `json:"password" required:"true"`
		ConfirmPassword string `json:"confirmPassword" required:"true"`
	}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		httpHelper.InternalServerError(w)
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
	uuid, err := crypto.CreateUUID()
	if err != nil {
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	hash, err := crypto.GenerateHash(data.Password)
	if err != nil {
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
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
	err = m.UserRepo.CreateUser(uuid, data.Login, hash, data.Email)
	if err != nil {
		if err.Error() == "user elready exists" {
			httpHelper.ConflictError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, http.StatusOK, models.DefaultMessage{
		Message: "User registered :)",
	})
}
