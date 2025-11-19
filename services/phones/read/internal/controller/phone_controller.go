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

	"parcial3/services/read/internal/repository"
	"parcial3/services/read/internal/service"
)

func NewHandler() http.Handler {
	// variables de entorno esperadas
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = "27017"
	}
	dbname := os.Getenv("MONGO_INITDB_DATABASE")
	if dbname == "" {
		dbname = "parcialdb"
	}

	if user == "" || pass == "" || host == "" {
		// fail-fast para detectar configuración faltante
		panic("missing MongoDB environment variables for read service")
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", user, pass, host, port)

	clientOpts := options.Client().ApplyURI(mongoURI)
	// conectar con timeout
	ctxConn, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctxConn, clientOpts)
	if err != nil {
		panic(fmt.Sprintf("mongo connect error: %v", err))
	}
	
	// ping
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelPing()
	if err := client.Ping(ctxPing, nil); err != nil {
		panic(fmt.Sprintf("mongo ping error: %v", err))
	}

	db := client.Database(dbname)
	repo := repository.NewPhoneRepository(db)
	svc := service.NewPhoneService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Read service OK"))
	})

	mux.HandleFunc("/phones", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		phones, err := svc.GetAll(r.Context())
		if err != nil {
			http.Error(w, "error reading phones: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(phones)
	})

	return mux
}
