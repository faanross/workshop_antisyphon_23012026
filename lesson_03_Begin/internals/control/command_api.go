package control

import "encoding/json"

// TODO: Add a Validator field to the struct in validCommands
// The Validator should be of type CommandValidator (defined below)
// Registry of valid commands with their validators and processors
var validCommands = map[string]struct {
	Validator CommandValidator
}{
	"shellcode": {
		// TODO: Assign validateShellcodeCommand as the Validator
		Validator: validateShellcodeCommand,
	},
}

// TODO: Define CommandValidator as a function type
// Hint: It takes json.RawMessage and returns error
// This allows us to validate command-specific arguments
type CommandValidator func(json.RawMessage) error
