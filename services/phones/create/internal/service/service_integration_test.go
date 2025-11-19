package service

import (
	"context"
	"testing"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertMongo(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	coll := client.Database("parcialdb").Collection("phones")

	_, err = coll.InsertOne(context.Background(), map[string]string{"name": "Test", "type": "integration"})
	if err != nil {
		t.Fatalf("failed insert: %v", err)
	}
}
