package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/camilo/parcial3/services/create/internal/repository"
	"github.com/camilo/parcial3/services/create/internal/service"
)

type Item struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func NewHandler() http.Handler {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	dbname := os.Getenv("MONGO_INITDB_DATABASE")

	if user == "" || pass == "" || host == "" {
		panic("missing MongoDB environment variables")
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=admin", user, pass, host)

	clientOpts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		panic(fmt.Sprintf("error connecting to MongoDB: %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("cannot ping MongoDB: %v", err))
	}

	db := client.Database(dbname)
	repo := repository.NewItemRepository(db)
	svc := &service.ItemService{Repo: repo}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Create service OK"))
	})

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := svc.Create(r.Context(), item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	})

	return mux
}
