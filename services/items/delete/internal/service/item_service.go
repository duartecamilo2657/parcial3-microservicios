package service

import (
	"context"

	"github.com/camilo/parcial3/services/delete/internal/repository"
)

type ItemService struct {
	repo *repository.DeleteRepository
}

func NewItemService(repo *repository.DeleteRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) DeleteItem(ctx context.Context, id interface{}) error {
	_, err := s.repo.DeleteByID(ctx, id)
	return err
}
