package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Phone struct {
	ID    interface{} `bson:"_id,omitempty" json:"_id"`
	Brand  string      `bson:"brand" json:"brand"`
	Price int         `bson:"price" json:"price"`
}

type PhoneRepository struct {
	collection *mongo.Collection
}

func NewPhoneRepository(db *mongo.Database) *PhoneRepository {
	return &PhoneRepository{
		collection: db.Collection("phones"),
	}
}

func (r *PhoneRepository) FindAll(ctx context.Context) ([]Phone, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var phones []Phone
	if err := cursor.All(ctx, &phones); err != nil {
		return nil, err
	}

	if phones == nil {
		phones = []Phone{}
	}

	return phones, nil
}
