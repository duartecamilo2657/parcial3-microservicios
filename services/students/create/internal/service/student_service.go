package service

import (
	"context"
	"parcial3/services/create/internal/repository"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func (s *StudentService) Create(ctx context.Context, student interface{}) error {
	return s.Repo.Create(ctx, student)
}
