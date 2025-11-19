package service

import (
	"context"

	"parcial3/services/read/internal/repository"
)

type PhoneService struct {
	repo *repository.PhoneRepository
}

func NewPhoneService(repo *repository.PhoneRepository) *PhoneService {
	return &PhoneService{repo: repo}
}

func (s *PhoneService) GetAll(ctx context.Context) ([]repository.Phone, error) {
	return s.repo.FindAll(ctx)
}
