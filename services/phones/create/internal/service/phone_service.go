package service

import (
	"context"
	"parcial3/services/create/internal/repository"
)

type PhoneService struct {
	Repo *repository.PhoneRepository
}

func (s *PhoneService) Create(ctx context.Context, phone interface{}) error {
	return s.Repo.Create(ctx, phone)
}
