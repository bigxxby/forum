package post

import (
	"database/sql"
	"forum/pkg/httpHelper"
	"net/http"
)

func (c *PostController) GET_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	postId := httpHelper.GetPostIdFromPath(r.URL.Path)
	if postId == -1 {
		httpHelper.BadRequestError(w)
		return
	}
	post, err := c.PostService.GetPostById(postId)
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
	}
	httpHelper.WriteJson(w, 200, post)
}
