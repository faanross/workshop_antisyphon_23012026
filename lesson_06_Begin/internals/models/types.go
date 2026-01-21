package models

import "encoding/json"

// CommandClient represents a command with its arguments as sent by Client
type CommandClient struct {
	Command   string          `json:"command"`
	Arguments json.RawMessage `json:"data,omitempty"`
}

// TODO: Define ServerResponse struct to represent a response from the server to the agent
// This tells the agent whether there's a job to execute
// Hint: It should have:
//   - Job (bool) with json tag "job" - indicates if there's a command
//   - JobID (string) with json tag "job_id,omitempty" - unique job identifier
//   - Command (string) with json tag "command,omitempty" - the command name
//   - Arguments (json.RawMessage) with json tag "data,omitempty" - command args
type ServerResponse struct {
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
