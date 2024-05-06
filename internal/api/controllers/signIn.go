package controllers

import (
	"database/sql"
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"io"
	"log"
	"net/http"
)

func (m *Manager) POST_SignIn(w http.ResponseWriter, r *http.Request) {
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
		httpHelper.InternalServerError(w)
		return
	}
	//////////////////////VALIDATION
	if data.Email == "" || data.Password == "" {
		httpHelper.BadRequestError(w)
		return
	}
	//////////////////////VALIDATION
	uuid, err := m.UserRepo.AuthUser(data.Email, data.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}

	httpHelper.WriteJson(w, http.StatusOK, models.AuthenticationMessage{
		Message: "Congratz, authorised :)",
		UUID:    uuid,
	})
}
