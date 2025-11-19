package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	Collection *mongo.Collection
}

func NewStudentRepository(db *mongo.Database) *StudentRepository {
	return &StudentRepository{Collection: db.Collection("students")}
}

func (r *StudentRepository) Create(ctx context.Context, student interface{}) error {
	_, err := r.Collection.InsertOne(ctx, student)
	return err
}
