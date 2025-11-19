package main

import (
	"log"
	"net/http"
	"os"

	"parcial3/services/read/internal/controller"
)

func main() {
	port := os.Getenv("READ_PORT")
	if port == "" {
		port = "8082"
	}

	handler := controller.NewHandler()
	log.Printf("Read service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("read service failed: %v", err)
	}
}
