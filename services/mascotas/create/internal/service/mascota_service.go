package service

import (
	"context"
	"parcial3/services/create/internal/repository"
)

type MascotaService struct {
	Repo *repository.MascotaRepository
}

func (s *MascotaService) Create(ctx context.Context, mascota interface{}) error {
	return s.Repo.Create(ctx, mascota)
}