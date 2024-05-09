package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"onePunchAkashMonitoring/pkg/node"
)

// NewRouter creates and returns a new router with all the configured routes.
func NewRouter(n *node.Node) *mux.Router {
	router := mux.NewRouter()

	// Endpoint for health checks
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Implement the logic to return the node's health status.
		// This could be as simple as an HTTP 200 OK if the node is running.
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	// Endpoint to receive updates or commands
	router.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		// Implement the logic to handle updates or commands sent to the node.
		w.WriteHeader(http.StatusAccepted)
	}).Methods("POST")

	// Additional endpoints can be defined here
	return router
}
