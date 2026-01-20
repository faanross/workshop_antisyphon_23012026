package control

import (
	"encoding/json"
	"log"
	"sync"
	"workshop3_dev/internals/models"
)

// Registry of valid commands with their validators and processors
var validCommands = map[string]struct {
	Validator CommandValidator
	Processor CommandProcessor
}{
	"shellcode": {
		Validator: validateShellcodeCommand,
		Processor: processShellcodeCommand,
	},
}

// CommandValidator validates command-specific arguments
type CommandValidator func(json.RawMessage) error

// CommandProcessor processes command-specific arguments
type CommandProcessor func(json.RawMessage) (json.RawMessage, error)

// TODO: Define CommandQueue struct to store commands ready for agent pickup
// Hint: It should have:
//   - PendingCommands: slice of models.CommandClient
//   - mu: sync.Mutex for thread-safe access
type CommandQueue struct {
	PendingCommands []models.CommandClient
	mu              sync.Mutex
}

// TODO: Create a global AgentCommands variable of type CommandQueue
// This is where validated/processed commands wait for agent pickup
// Hint: Initialize PendingCommands with make([]models.CommandClient, 0)
var AgentCommands = CommandQueue{
	PendingCommands: make([]models.CommandClient, 0),
}

// TODO: Implement addCommand method to add a validated command to the queue
// Remember to use mutex for thread safety!
// Hint: Lock mutex, append to slice, unlock
func (cq *CommandQueue) addCommand(command models.CommandClient) {
	// TODO: Lock the mutex
	cq.mu.Lock()
	// TODO: Defer unlock
	defer cq.mu.Unlock()

	// TODO: Append command to PendingCommands
	cq.PendingCommands = append(cq.PendingCommands, command)
	log.Printf("QUEUED: %s", command.Command)
}
