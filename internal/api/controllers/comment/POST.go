package comment

import (
	"encoding/json"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"io"
	"log"
	"net/http"
)

func (c *CommentController) POST_Comment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	postId := httpHelper.GetIdFromString(r.PathValue("id"))
	if postId == -1 {
		httpHelper.BadRequestError(w)
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
	}{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err.Error())
		httpHelper.BadRequestError(w)
		return
	}

	commentId, err := c.CommentService.PostComment(userIdNum, postId, data.Content)
	if err != nil {
		if err.Error() == "comment is not valid" {

			httpHelper.BadRequestError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, models.CreationMessage{
		Message: "Comment Created :)",
		Id:      commentId,
	})

}
