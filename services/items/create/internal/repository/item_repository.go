package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	Collection *mongo.Collection
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	return &ItemRepository{Collection: db.Collection("items")}
}

func (r *ItemRepository) Create(ctx context.Context, item interface{}) error {
	_, err := r.Collection.InsertOne(ctx, item)
	return err
}
