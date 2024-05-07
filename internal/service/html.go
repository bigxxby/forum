package service

import "forum/internal/repository"

type HTMLService struct {
	HTMLRepository *repository.HTMLRepository
}

func NewHTMLService(repo *repository.HTMLRepository) *HTMLService {
	return &HTMLService{
		HTMLRepository: repo,
	}
}
