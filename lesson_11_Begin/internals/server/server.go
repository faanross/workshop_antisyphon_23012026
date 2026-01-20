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

	// TODO: Add POST endpoint for receiving results from agent
	// Hint: r.Post("/results", ResultHandler)

	// Create the HTTP server
	server.server = &http.Server{
		Addr:    server.addr,
		Handler: r,
	}

	// Start the server
	return server.server.ListenAndServeTLS(server.tlsCert, server.tlsKey)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	var response models.ServerResponse

	// Check for pending commands
	cmd, exists := control.AgentCommands.GetCommand()
	if exists {
		log.Printf("Sending command to agent: %s\n", cmd.Command)
		response.Job = true
		response.Command = cmd.Command
		response.Arguments = cmd.Arguments
		response.JobID = fmt.Sprintf("job_%06d", rand.Intn(1000000))
		log.Printf("Job ID: %s\n", response.JobID)
	} else {
		log.Printf("No commands in queue")
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
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

// TODO: Implement ResultHandler to receive and display results from the Agent
// Hint: This should:
// 1. Decode the request body into models.AgentTaskResult
// 2. Unmarshal the CommandResult field to get the message string
// 3. Log success or failure based on result.Success field
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	// TODO: Decode the incoming result into models.AgentTaskResult
	// Hint: var result models.AgentTaskResult
	// json.NewDecoder(r.Body).Decode(&result)

	// TODO: Unmarshal the CommandResult to get the actual message string
	// Hint: var messageStr string
	// json.Unmarshal(result.CommandResult, &messageStr)

	// TODO: Log success or failure based on result.Success
	// Hint: if !result.Success { log failure } else { log success }
}
