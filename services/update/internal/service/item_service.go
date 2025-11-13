package service

import (
	"context"

	"github.com/camilo/parcial3/services/update/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type ItemService struct {
	repo *repository.UpdateRepository
}

func NewItemService(repo *repository.UpdateRepository) *ItemService {
	return &ItemService{repo: repo}
}

// UpdateItem recibe id (primitive.ObjectID o string) y un mapa con campos a actualizar
func (s *ItemService) UpdateItem(ctx context.Context, id interface{}, name string, value int) error {
	update := bson.M{}
	if name != "" {
		update["name"] = name
	}
	update["value"] = value
	_, err := s.repo.UpdateByID(ctx, id, update)
	return err
}
