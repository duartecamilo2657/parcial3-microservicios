package service

import (
	"context"

	"parcial3/services/read/internal/repository"
)

type MascotaService struct {
	repo *repository.MascotaRepository
}

func NewMascotaService(repo *repository.MascotaRepository) *MascotaService {
	return &MascotaService{repo: repo}
}

func (s *MascotaService) GetAll(ctx context.Context) ([]repository.Mascota, error) {
	return s.repo.FindAll(ctx)
}
