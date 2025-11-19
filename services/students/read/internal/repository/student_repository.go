package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	ID   interface{} `bson:"_id,omitempty" json:"_id"`
	Name string      `bson:"name" json:"name"`
	Age  int         `bson:"age" json:"age"`
}

type StudentRepository struct {
	collection *mongo.Collection
}

func NewStudentRepository(db *mongo.Database) *StudentRepository {
	return &StudentRepository{
		collection: db.Collection("students"),
	}
}

func (r *StudentRepository) FindAll(ctx context.Context) ([]Student, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []Student
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	if students == nil {
		students = []Student{}
	}

	return students, nil
}
