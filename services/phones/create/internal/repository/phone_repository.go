package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneRepository struct {
	Collection *mongo.Collection
}

func NewPhoneRepository(db *mongo.Database) *PhoneRepository {
	return &PhoneRepository{Collection: db.Collection("phones")}
}

func (r *PhoneRepository) Create(ctx context.Context, phone interface{}) error {
	_, err := r.Collection.InsertOne(ctx, phone)
	return err
}
