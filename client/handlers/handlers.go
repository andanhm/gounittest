package handlers

import (
	"io"
	"net/http"
)

// HealthCheck returns the status 200 OK
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `Ok`)
}
