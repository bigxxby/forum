package comment

import (
	"forum/pkg/httpHelper"
	"log"
	"net/http"
)

func (c *CommentController) GET_Comments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	postId := httpHelper.GetIdFromString(r.PathValue("postId"))
	if postId == -1 {
		httpHelper.BadRequestError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, _ := userId.(int)
	comments, err := c.CommentService.GetAllCommentsOfAPost(postId, userIdNum)
	if err != nil {
		httpHelper.WriteJson(w, 404, comments)
		return
	}
	httpHelper.WriteJson(w, 404, comments)

}
func (c *CommentController) GET_LikedComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, _ := userId.(int)
	comments, err := c.CommentService.GetAllLikedComments(userIdNum)
	if err != nil {
		log.Println(err.Error())
		httpHelper.WriteJson(w, 404, comments)
		return
	}
	httpHelper.WriteJson(w, 404, comments)

}
