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

type CommandQueue struct {
	//   - PendingCommands: slice of models.CommandClient
	//   - mu: sync.Mutex for thread-safe access
}

// TODO: Create a global AgentCommands variable of type CommandQueue
var AgentCommands = CommandQueue{}

// TODO: Implement addCommand method to add a validated command to the queue
// Remember to use mutex for thread safety!
// Hint: Lock mutex, append to slice, unlock
func (cq *CommandQueue) addCommand(command models.CommandClient) {

	cq.mu.Lock()
	defer cq.mu.Unlock()

	// TODO: Append command to PendingCommands
	log.Printf("QUEUED: %s", command.Command)
}
