package core

import (
	"log"
	"net/http"
)

// StartHealthCheckServer starts the HTTP server for health check
func StartHealthCheckServer(port string) {
	http.HandleFunc("/health", healthCheckHandler)
	log.Println("Starting Health Check Server on port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Health Check Server failed: %s", err)
	}
}

// healthCheckHandler processes status check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Checking the service status
	//ToDo Additional checks can be added here

	// If all checks are successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
