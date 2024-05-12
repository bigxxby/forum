package post

import (
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
		Content string `json:"content"`
		Title   string `json:"title"`
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
	if data.Content == "" || data.Title == "" {
		httpHelper.BadRequestError(w)
		log.Println("2")

		return
	}
	if !validation.IsValidPost(data.Title, data.Content) {
		httpHelper.BadRequestError(w)
		log.Println("1")
		return
	}
	//VALIDATION

	postId, err := c.PostService.CreatePost(userIdNum, data.Title, data.Content)
	if err != nil {
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, http.StatusOK, models.PostCreationMessage{
		Message: "Post created :)",
		PostId:  postId,
	})

}
