package service

import (
	"context"
	"github.com/camilo/parcial3/services/create/internal/repository"
)

type ItemService struct {
	Repo *repository.ItemRepository
}

func (s *ItemService) Create(ctx context.Context, item interface{}) error {
	return s.Repo.Create(ctx, item)
}
