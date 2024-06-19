package post

import (
	"database/sql"
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"forum/pkg/validation"
	"io"
	"log"
	"net/http"
)

func (c *PostController) POST_PostPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/posts" {
		httpHelper.NotFoundError(w)
		return

	}
	userId := r.Context().Value("userId")
	userIdNum, ok := userId.(int)
	if !ok {
		httpHelper.Unauthorised(w)
		return
	}
	data := struct {
		Content    string   `json:"content"`
		Title      string   `json:"title"`
		Categories []string `json:"categories"`
	}{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err.Error())
		httpHelper.InternalServerError(w)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err.Error())
		httpHelper.BadRequestError(w)
		return
	}
	//VALIDATION
	if userId == "" {
		httpHelper.Unauthorised(w)
		return
	}
	if data.Content == "" || data.Title == "" || len(data.Categories) == 0 {
		httpHelper.BadRequestError(w)
		return
	}
	if !validation.IsValidPost(data.Title, data.Content) {
		httpHelper.BadRequestError(w)
		return
	}
	//VALIDATION

	postId, err := c.PostService.CreatePost(userIdNum, data.Title, data.Content, data.Categories)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		if err == models.ErrBadRequest {
			httpHelper.BadRequestError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, http.StatusOK, models.CreationMessage{
		Message: "Post created :)",
		Id:      postId,
	})

}
