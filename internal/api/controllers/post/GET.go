package post

import (
	"database/sql"
	"forum/pkg/httpHelper"
	"log"
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
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, post)
}
func (c *PostController) GET_posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/posts/all" {
		httpHelper.NotFoundError(w)
		return
	}
	posts, err := c.PostService.GetAllPostsByCreationTime()
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)
		return
	}
	httpHelper.WriteJson(w, 200, posts)

}
