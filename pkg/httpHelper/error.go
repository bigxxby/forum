package httpHelper

import (
	"forum/internal/models"
	"net/http"
)

func MethodNotAllowedError(w http.ResponseWriter) {
	WriteJson(w, http.StatusMethodNotAllowed, models.DefaultMessage{
		Message: "Method not allowed, sorry :}",
	})
}

func NotFoundError(w http.ResponseWriter) {
	WriteJson(w, http.StatusNotFound, models.DefaultMessage{
		Message: "Nothing found... :|",
	})
}

func InternalServerError(w http.ResponseWriter) {
	WriteJson(w, http.StatusInternalServerError, models.DefaultMessage{
		Message: "Whoops, something broke :(",
	})
}
func Unauthorised(w http.ResponseWriter) {
	WriteJson(w, http.StatusUnauthorized, models.DefaultMessage{
		Message: "Permisson denied - GET OUT OF HERE >:(",
	})
}

func BadRequestError(w http.ResponseWriter) {
	WriteJson(w, http.StatusBadRequest, models.DefaultMessage{
		Message: "Bad request, buddy :P",
	})
}
