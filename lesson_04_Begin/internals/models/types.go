package models

import "encoding/json"

// CommandClient represents a command with its arguments as sent by Client
type CommandClient struct {
	Command   string          `json:"command"`
	Arguments json.RawMessage `json:"data,omitempty"`
}

// ShellcodeArgsClient contains the command-specific arguments for Shellcode Loader as sent by Client
type ShellcodeArgsClient struct {
	FilePath   string `json:"file_path"`
	ExportName string `json:"export_name"`
}

// TODO: Define ShellcodeArgsAgent struct for arguments sent TO the Agent
// This is different from ShellcodeArgsClient because:
// - Instead of FilePath, we send the actual shellcode as base64
// Hint: It should have two fields:
//   - ShellcodeBase64 (string) with json tag "shellcode_base64"
//   - ExportName (string) with json tag "export_name"
type ShellcodeArgsAgent struct {
	// TODO: Add ShellcodeBase64 field (string)
	// TODO: Add ExportName field (string)
	ExportName string `json:"export_name"`
}
