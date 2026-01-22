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

// CommandQueue stores commands ready for agent pickup
type CommandQueue struct {
	PendingCommands []models.CommandClient
	mu              sync.Mutex
}

// AgentCommands is Global command queue
var AgentCommands = CommandQueue{
	PendingCommands: make([]models.CommandClient, 0),
}

// addCommand adds a validated command to the queue
func (cq *CommandQueue) addCommand(command models.CommandClient) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	cq.PendingCommands = append(cq.PendingCommands, command)
	log.Printf("QUEUED: %s", command.Command)
}

// TODO: Implement GetCommand to retrieve and remove the next command from queue
func (cq *CommandQueue) GetCommand() (models.CommandClient, bool) {

	cq.mu.Lock()
	defer cq.mu.Unlock()

	// TODO: if there is nothing in queue (PendingCommands), return false

	// TODO: init variable cmd equal to value at index 0 in PendingCommands

	// TODO: Remove it from the queue (slice from index 1)

	log.Printf("DEQUEUED: Command '%s'", cmd.Command)

	return cmd, true
}
