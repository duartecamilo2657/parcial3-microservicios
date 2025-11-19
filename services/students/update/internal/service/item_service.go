package service

import (
	"context"

	"parcial3/services/update/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type StudentService struct {
	repo *repository.UpdateRepository
}

func NewStudentService(repo *repository.UpdateRepository) *StudentService {
	return &StudentService{repo: repo}
}

// UpdateStudent recibe id (primitive.ObjectID o string) y un mapa con campos a actualizar
func (s *StudentService) UpdateStudent(ctx context.Context, id interface{}, name string, age int) error {
	update := bson.M{}
	if name != "" {
		update["name"] = name
	}
	update["age"] = age
	_, err := s.repo.UpdateByID(ctx, id, update)
	return err
}
