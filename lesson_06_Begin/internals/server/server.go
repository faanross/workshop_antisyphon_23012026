package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"math/rand"
	"net/http"
	"time"
	"workshop3_dev/internals/control"
	"workshop3_dev/internals/models"
)

// Server implements the Server interface for HTTPS
type Server struct {
	addr    string
	server  *http.Server
	tlsCert string
	tlsKey  string
}

// NewServer creates a new HTTPS server
func NewServer(addr string) *Server {
	return &Server{
		addr:    addr,
		tlsCert: "./certs/server.crt",
		tlsKey:  "./certs/server.key",
	}
}

// Start implements Server.Start for HTTPS
func (server *Server) Start() error {
	// Create Chi router
	r := chi.NewRouter()

	// Define our GET endpoint
	r.Get("/", RootHandler)

	// Create the HTTP server
	server.server = &http.Server{
		Addr:    server.addr,
		Handler: r,
	}

	// Start the server
	return server.server.ListenAndServeTLS(server.tlsCert, server.tlsKey)
}

// TODO: Update RootHandler to return ServerResponse with any pending commands
func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	var response models.ServerResponse

	// TODO: Check for pending commands using control.AgentCommands.GetCommand()

	if exists {

		log.Printf("Sending command to agent: %s\n", cmd.Command)

		// If command exists, populate the response
		// TODO set response.Job to true
		// TODO set response.Command
		// TODO set response.Arguments

		response.JobID = fmt.Sprintf("job_%06d", rand.Intn(1000000))
		log.Printf("Job ID: %s\n", response.JobID)
	} else {
		log.Printf("No commands in queue")
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// TODO: Encode and send the ServerResponse
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// Stop implements Server.Stop for HTTPS
func (server *Server) Stop() error {
	// If there's no server, nothing to stop
	if server.server == nil {
		return nil
	}

	// Give the server 5 seconds to shut down gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.server.Shutdown(ctx)
}
