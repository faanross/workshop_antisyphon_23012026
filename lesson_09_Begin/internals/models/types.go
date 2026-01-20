package models

import "encoding/json"

// CommandClient represents a command with its arguments as sent by Client
type CommandClient struct {
	Command   string          `json:"command"`
	Arguments json.RawMessage `json:"data,omitempty"`
}

// ServerResponse represents a response from the server to the agent
type ServerResponse struct {
	Job       bool            `json:"job"`
	JobID     string          `json:"job_id,omitempty"`
	Command   string          `json:"command,omitempty"`
	Arguments json.RawMessage `json:"data,omitempty"`
}

type AgentTaskResult struct {
	JobID         string          `json:"job_id"`
	Success       bool            `json:"success"`
	CommandResult json.RawMessage `json:"command_result,omitempty"`
	Error         error           `json:"error,omitempty"`
}

// ShellcodeArgsClient contains the command-specific arguments for Shellcode Loader as sent by Client
type ShellcodeArgsClient struct {
	FilePath   string `json:"file_path"`
	ExportName string `json:"export_name"`
}

// ShellcodeArgsAgent contains the command-specific arguments for Shellcode Loader as sent to the Agent
type ShellcodeArgsAgent struct {
	ShellcodeBase64 string `json:"shellcode_base64"`
	ExportName      string `json:"export_name"`
}

// TODO: Define ShellcodeResult struct to hold the result of shellcode execution
// This is returned by the shellcode doer
// Hint: It should have:
//   - Message (string) with json tag "message"
type ShellcodeResult struct {
	Message string `json:"message"`
}
