package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MascotaRepository struct {
	Collection *mongo.Collection
}

func NewMascotaRepository(db *mongo.Database) *MascotaRepository {
	return &MascotaRepository{Collection: db.Collection("mascotas")}
}

func (r *MascotaRepository) Create(ctx context.Context, mascota interface{}) error {
	_, err := r.Collection.InsertOne(ctx, mascota)
	return err
}
