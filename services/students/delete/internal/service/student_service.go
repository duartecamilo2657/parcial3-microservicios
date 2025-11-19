package service

import (
	"context"

	"parcial3/services/delete/internal/repository"
)

type StudentService struct {
	repo *repository.DeleteRepository
}

func NewStudentService(repo *repository.DeleteRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) DeleteStudent(ctx context.Context, id interface{}) error {
	_, err := s.repo.DeleteByID(ctx, id)
	return err
}
