package repository

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFindAllIntegration(t *testing.T) {
	uri := "mongodb://admin:admin123@localhost:27017/?authSource=admin"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("new client err: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		t.Fatalf("connect err: %v", err)
	}
	db := client.Database("parcialdb")
	repo := NewPhoneRepository(db)
	_, err = repo.FindAll(context.Background())
	if err != nil && err != mongo.ErrNoDocuments {
		// no fallamos si no hay docs; sólo verificamos que la consulta corra
		t.Fatalf("find all err: %v", err)
	}
}
