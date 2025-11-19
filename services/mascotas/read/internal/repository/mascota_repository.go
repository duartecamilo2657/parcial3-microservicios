package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mascota struct {
	ID    interface{} `bson:"_id,omitempty" json:"_id"`
	Name  string      `bson:"name" json:"name"`
	Age int         `bson:"age" json:"age"`
}

type MascotaRepository struct {
	collection *mongo.Collection
}

func NewMascotaRepository(db *mongo.Database) *MascotaRepository {
	return &MascotaRepository{
		collection: db.Collection("mascotas"),
	}
}

func (r *MascotaRepository) FindAll(ctx context.Context) ([]Mascota, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var mascotas []Mascota
	if err := cursor.All(ctx, &mascotas); err != nil {
		return nil, err
	}

	if mascotas == nil {
		mascotas = []Mascota{}
	}

	return mascotas, nil
}
