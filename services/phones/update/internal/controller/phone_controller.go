package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"parcial3/services/update/internal/repository"
	"parcial3/services/update/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReqUpdate struct {
	Brand  string `json:"brand"`
	Price int    `json:"price"`
}

func NewHandler() http.Handler {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = "27017"
	}
	dbname := os.Getenv("MONGO_INITDB_DATABASE")

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", user, pass, host, port)
	clientOpts := options.Client().ApplyURI(mongoURI)
	ctxConn, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctxConn, clientOpts)
	if err != nil {
		panic(err)
	}
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelPing()
	if err := client.Ping(ctxPing, nil); err != nil {
		panic(err)
	}

	db := client.Database(dbname)
	repo := repository.NewUpdateRepository(db)
	svc := service.NewPhoneService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Update service OK"))
	})

	// PUT /phones/{id}
	mux.HandleFunc("/phones/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// extraer id de la ruta
		idHex := r.URL.Path[len("/phones/"):]

		objID, err := primitive.ObjectIDFromHex(idHex)
		var id interface{}
		if err != nil {
			// no es ObjectID hex? intentar como string - pero mongo uses ObjectID typically
			http.Error(w, "invalid id format", http.StatusBadRequest)
			return
		} else {
			id = objID
		}

		var payload ReqUpdate
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
			return
		}
		if err := svc.UpdatePhone(r.Context(), id, payload.Brand, payload.Price); err != nil {
			http.Error(w, "update error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"updated"}`))
	})

	return mux
}
