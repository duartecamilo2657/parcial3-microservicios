package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Item struct {
	ID    interface{} `bson:"_id,omitempty" json:"_id"`
	Name  string      `bson:"name" json:"name"`
	Value int         `bson:"value" json:"value"`
}

type ItemRepository struct {
	collection *mongo.Collection
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	return &ItemRepository{
		collection: db.Collection("items"),
	}
}

func (r *ItemRepository) FindAll(ctx context.Context) ([]Item, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []Item
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}
