package category

import "forum/internal/service/category"

type CategoryController struct {
	CategoryService *category.CategoryService
}

func NewCategoryController(service *category.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: service,
	}
}
