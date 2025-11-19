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

	"parcial3/services/create/internal/repository"
	"parcial3/services/create/internal/service"
)

type Phone struct {
	Brand string `json:"brand"`
	Price int    `json:"price"`
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
	repo := repository.NewPhoneRepository(db)
	svc := &service.PhoneService{Repo: repo}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Create service OK"))
	})

	mux.HandleFunc("/phones", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var phone Phone
		if err := json.NewDecoder(r.Body).Decode(&phone); err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		// ===============================
		// VALIDACIONES DE REGLAS DE NEGOCIO
		// ===============================

		// Regla 1: brand NO puede ser vacío
		if phone.Brand == "" {
			http.Error(w, "brand cannot be empty", http.StatusBadRequest)
			return
		}

		// Regla 2: price debe ser mayor que 0
		if phone.Price <= 0 {
			http.Error(w, "price must be greater than 0", http.StatusBadRequest)
			return
		}

		// ===============================
		// SI TODO ESTÁ BIEN → GUARDAR
		// ===============================

		if err := svc.Create(r.Context(), phone); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(phone)
	})

	return mux
}
