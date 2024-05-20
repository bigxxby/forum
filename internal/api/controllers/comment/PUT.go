package comment

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

func (c *CommentController) PUT_EditComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	commentId := httpHelper.GetIdFromString(r.PathValue("id"))
	if commentId == -1 {
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
	///VALIDATION
	if !validation.IsValidComment(data.Content) {
		httpHelper.BadRequestError(w)
		return
	}
	///VALIDATION

	err = c.CommentService.EditComment(data.Content, commentId, userIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.Unauthorised(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, models.DefaultMessage{
		Message: "Comment edited :)",
	})

}
