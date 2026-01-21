package control

import "encoding/json"

// Registry of valid commands with their validators and processors
var validCommands = map[string]struct {
	Validator CommandValidator
	// TODO: Add Processor field of type CommandProcessor

}{
	"shellcode": {
		Validator: validateShellcodeCommand,
		// TODO: Assign processShellcodeCommand as the Processor

	},
}

// CommandValidator validates command-specific arguments
type CommandValidator func(json.RawMessage) error

// TODO: Define CommandProcessor as a function type
// Hint: It takes json.RawMessage and returns (json.RawMessage, error)
// This transforms client args into agent args (e.g., reads file and encodes to base64)
