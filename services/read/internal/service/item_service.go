package service

import (
	"context"

	"github.com/camilo/parcial3/services/read/internal/repository"
)

type ItemService struct {
	repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) GetAll(ctx context.Context) ([]repository.Item, error) {
	return s.repo.FindAll(ctx)
}
