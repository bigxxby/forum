package middlewares

import (
	"bytes"
	"context"
	"forum/internal/service/user"
	"io"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc, service *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "userId", "")
		cookie, err := r.Cookie("uuid")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		uuid := cookie.Value
		user, err := service.UserRepository.GetUserByUUID(uuid)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		ctx = context.WithValue(r.Context(), "userId", user.ID)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
