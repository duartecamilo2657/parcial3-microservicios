package main

import (
	"log"
	"net/http"
	"os"

	"parcial3/services/update/internal/controller"
)

func main() {
	port := os.Getenv("UPDATE_PORT")
	if port == "" {
		port = "8083"
	}
	handler := controller.NewHandler()
	log.Printf("Update service listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
