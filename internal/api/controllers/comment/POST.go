package comment

import (
	"database/sql"
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
func (c *CommentController) POST_Like(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	err := c.CommentService.LikeComment(userIdNum, commentId)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		if err.Error() == "comment already liked" {
			httpHelper.ConflictError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, models.DefaultMessage{
		Message: "Liked :)",
	})
}
func (c *CommentController) POST_UnLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	err := c.CommentService.UnLikeComment(userIdNum, commentId)

	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		if err.Error() == "comment not liked" {
			httpHelper.ConflictError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, models.DefaultMessage{
		Message: "UnLiked :(",
	})
}
func (c *CommentController) POST_Reply(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	parentId := httpHelper.GetIdFromString(r.PathValue("parentId"))
	if parentId == -1 {
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

	commentId, err := c.CommentService.ReplyToComment(parentId, userIdNum, data.Content)

	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		if err.Error() == "comment is not valid" {
			httpHelper.BadRequestError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, models.CreationMessage{
		Message: "Replied :)",
		Id:      commentId,
	})
}
