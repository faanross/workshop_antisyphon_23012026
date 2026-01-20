package models

// TODO: Define CommandClient struct to represent a command with its arguments as sent by Client
// Hint: It should have two fields:
//   - Command (string) with json tag "command"
//   - Arguments (json.RawMessage) with json tag "data,omitempty"
type CommandClient struct {
	// TODO: Add Command field (string)
	// TODO: Add Arguments field (json.RawMessage)
}

// TODO: Define ShellcodeArgsClient struct for shellcode command arguments
// Hint: It should have two fields:
//   - FilePath (string) with json tag "file_path"
//   - ExportName (string) with json tag "export_name"
type ShellcodeArgsClient struct {
	// TODO: Add FilePath field (string)
	// TODO: Add ExportName field (string)
}
