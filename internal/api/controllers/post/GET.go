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
	postId := httpHelper.GetIdFromString(r.PathValue("id"))
	if postId == -1 {
		httpHelper.BadRequestError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, _ := userId.(int)
	post, err := c.PostService.GetPostById(postId, userIdNum)
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
	userId := r.Context().Value("userId")
	userIdNum, _ := userId.(int)

	posts, err := c.PostService.GetAllPostsByCreationTime(userIdNum)
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
func (c *PostController) GET_likedPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/posts/liked" {
		httpHelper.NotFoundError(w)
		return
	}
	userId := r.Context().Value("userId")
	userIdNum, _ := userId.(int)

	posts, err := c.PostService.GetAllLikedPosts(userIdNum)
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
