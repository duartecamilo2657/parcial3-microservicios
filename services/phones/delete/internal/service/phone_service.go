package service

import (
	"context"

	"parcial3/services/delete/internal/repository"
)

type PhoneService struct {
	repo *repository.DeleteRepository
}

func NewPhoneService(repo *repository.DeleteRepository) *PhoneService {
	return &PhoneService{repo: repo}
}

func (s *PhoneService) DeletePhone(ctx context.Context, id interface{}) error {
	_, err := s.repo.DeleteByID(ctx, id)
	return err
}
