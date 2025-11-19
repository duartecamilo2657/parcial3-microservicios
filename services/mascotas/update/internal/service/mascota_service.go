package service

import (
	"context"

	"parcial3/services/update/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type MascotaService struct {
	repo *repository.UpdateRepository
}

func NewMascotaService(repo *repository.UpdateRepository) *MascotaService {
	return &MascotaService{repo: repo}
}


func (s *MascotaService) UpdateMascota(ctx context.Context, id interface{}, name string, age int) error {
	update := bson.M{}

	if name != "" {
		update["name"] = name
	}
	update["age"] = age
	_, err := s.repo.UpdateByID(ctx, id, update)
	return err
}
