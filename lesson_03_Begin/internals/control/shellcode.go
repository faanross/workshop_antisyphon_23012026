package control

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"workshop3_dev/internals/models"
)

// TODO: Implement validateShellcodeCommand to validate "shellcode" command arguments
// This function should:
// 1. Check if rawArgs is empty
// 2. Unmarshal rawArgs into models.ShellcodeArgsClient
// 3. Validate that FilePath is not empty
// 4. Validate that ExportName is not empty
// 5. Check if the file exists using os.Stat
func validateShellcodeCommand(rawArgs json.RawMessage) error {
	// TODO: Check if arguments are empty
	if len(rawArgs) == 0 {
		return fmt.Errorf("shellcode command requires arguments")
	}

	// TODO: Create variable to hold unmarshalled arguments
	var args models.ShellcodeArgsClient

	// TODO: Unmarshal the raw JSON into the args struct
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	// TODO: Validate FilePath is not empty
	if args.FilePath == "" {
		return fmt.Errorf("file_path is required")
	}

	// TODO: Validate ExportName is not empty
	if args.ExportName == "" {
		return fmt.Errorf("export_name is required")
	}

	// TODO: Check if file exists using os.Stat
	// Hint: os.IsNotExist(err) checks if file doesn't exist
	if _, err := os.Stat(args.FilePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", args.FilePath)
	}

	log.Printf("Validation passed: file_path=%s, export_name=%s", args.FilePath, args.ExportName)

	return nil
}
