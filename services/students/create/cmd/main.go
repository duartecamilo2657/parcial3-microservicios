package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"parcial3/services/create/internal/controller"
)

func main() {
	port := os.Getenv("CREATE_PORT")
	if port == "" {
		port = "8081"
	}
	handler := controller.NewHandler()
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Create service listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
