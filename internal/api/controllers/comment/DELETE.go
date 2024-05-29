package comment

import (
	"database/sql"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"log"
	"net/http"
)

func (c *CommentController) DELETE_Comment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	commentId := httpHelper.GetIdFromString(r.URL.Query().Get("valueId"))
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

	err := c.CommentService.DeleteComment(userIdNum, commentId)
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
		Message: "Comment Deleted :(",
	})
}
