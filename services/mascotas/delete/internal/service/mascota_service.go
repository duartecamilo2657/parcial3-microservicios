package service

import (
	"context"

	"parcial3/services/delete/internal/repository"
)

type MascotaService struct {
	repo *repository.DeleteRepository
}

func NewMascotaService(repo *repository.DeleteRepository) *MascotaService {
	return &MascotaService{repo: repo}
}

func (s *MascotaService) DeleteMascota(ctx context.Context, id interface{}) error {
	_, err := s.repo.DeleteByID(ctx, id)
	return err
}

