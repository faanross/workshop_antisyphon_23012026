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

	// TODO: Create a variable of type models.CommandClient to receive the command

	if err := json.NewDecoder(r.Body).Decode(&cmdClient); err != nil {
		log.Printf("ERROR: Failed to decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error decoding JSON")
		return
	}

	// TODO: Log the received command
	// Hint: Use fmt.Sprintf to create a message like "Received command: %s"
	log.Printf(commandReceived)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commandReceived)

}
