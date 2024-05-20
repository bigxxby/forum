package category

import (
	"database/sql"
	"forum/pkg/httpHelper"
	"log"
	"net/http"
)

func (c *CategoryController) GET_categories(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		httpHelper.MethodNotAllowedError(w)
		return
	}
	if r.URL.Path != "/api/categories" {
		httpHelper.NotFoundError(w)
		return
	}
	categories, err := c.CategoryService.GetAllCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			httpHelper.NotFoundError(w)
			return
		}
		log.Println(err.Error())
		httpHelper.InternalServerError(w)

		return

	}
	if categories == nil {
		httpHelper.NotFoundError(w)
		return
	}
	httpHelper.WriteJson(w, 200, categories)
}
