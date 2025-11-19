package service

import (
	"context"

	"parcial3/services/read/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetAll(ctx context.Context) ([]repository.Student, error) {
	return s.repo.FindAll(ctx)
}
