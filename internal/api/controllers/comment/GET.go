package comment

import (
	"forum/pkg/httpHelper"
	"net/http"
)

func (c *CommentController) GET_Comments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	postId := httpHelper.GetIdFromString(r.PathValue("id"))
	if postId == -1 {
		httpHelper.BadRequestError(w)
		return
	}
	comments, err := c.CommentService.GetAllCommentsOfAPost(postId)
	if err != nil {
		httpHelper.WriteJson(w, 404, comments)
		return
	}
	httpHelper.WriteJson(w, 404, comments)

}
