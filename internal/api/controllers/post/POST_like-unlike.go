package post

import (
	"database/sql"
	"forum/internal/models"
	"forum/pkg/httpHelper"
	"log"
	"net/http"
)

func (c *PostController) POST_Like(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	postId := r.PathValue("id")
	postIdNum := httpHelper.GetIdFromString(postId)
	if postIdNum == -1 {
		httpHelper.BadRequestError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, ok := userId.(int)
	if !ok {
		httpHelper.Unauthorised(w)
		return
	}

	err := c.PostService.LikePost(userIdNum, postIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		if err.Error() == "post already liked" {
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

////////

// //
func (c *PostController) POST_Unlike(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	postId := r.PathValue("id")
	postIdNum := httpHelper.GetIdFromString(postId)
	if postIdNum == -1 {
		httpHelper.BadRequestError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, ok := userId.(int)
	if !ok {
		httpHelper.Unauthorised(w)
		return
	}

	err := c.PostService.UnLikePost(userIdNum, postIdNum)
	if err != nil {
		if err.Error() == "post not liked" {
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
