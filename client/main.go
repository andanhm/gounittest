package main

import (
	"net/http"

	"github.com/andanhm/gounittest/client/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthCheck)
	http.ListenAndServe(":8080", nil)
}
