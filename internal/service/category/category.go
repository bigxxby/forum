package category

import "forum/internal/repository/category"

type CategoryService struct {
	CategoryRepository *category.CategoryRepository
}

func NewCategoryService(repo *category.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: repo,
	}
}
func (s *CategoryService) GetAllCategories() (map[string]int, error) {
	categories, err := s.CategoryRepository.SELECT_categories()
	if err != nil {
		return nil, err
	}
	return categories, nil

}
