package control

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"workshop3_dev/internals/models"
)

func StartControlAPI() {
	// Create Chi router
	r := chi.NewRouter()

	// TODO: Change this from GET to POST and change endpoint from "/dummy" to "/command"
	r.Get("/dummy", dummyHandler)

	log.Println("Starting Control API on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Printf("Control API error: %v", err)
		}
	}()
}

// TODO: Implement commandHandler to receive and parse commands from the client
// This replaces the old dummyHandler
func commandHandler(w http.ResponseWriter, r *http.Request) {

	// TODO create cmdClient of our new type CommandClient

	// The first thing we need to do is unmarshal the request body into the custom type
	if err := json.NewDecoder(r.Body).Decode(&cmdClient); err != nil {
		log.Printf("ERROR: Failed to decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error decoding JSON")
		return
	}

	// Visually confirm we get the command we expected
	// TODO Create a variable commandReceived, compose using Sprintf and cmdClient.Command
	log.Printf(commandReceived)

	// Confirm on the client side command was received
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commandReceived)
}
