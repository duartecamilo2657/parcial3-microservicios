package service

import (
	"context"

	"parcial3/services/update/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type PhoneService struct {
	repo *repository.UpdateRepository
}

func NewPhoneService(repo *repository.UpdateRepository) *PhoneService {
	return &PhoneService{repo: repo}
}

// UpdatePhone recibe id (primitive.ObjectID o string) y un mapa con campos a actualizar
func (s *PhoneService) UpdatePhone(ctx context.Context, id interface{}, brand string, price int) error {
	update := bson.M{}
	if brand != "" {
		update["brand"] = brand
	}
	update["price"] = price
	_, err := s.repo.UpdateByID(ctx, id, update)
	return err
}
