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

	if len(rawArgs) == 0 {
		return fmt.Errorf("shellcode command requires arguments")
	}

	// TODO: Create variable to hold unmarshalled arguments

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	if args.FilePath == "" {
		return fmt.Errorf("file_path is required")
	}

	// TODO: Validate ExportName is not empty

	if _, err := os.Stat(args.FilePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", args.FilePath)
	}

	log.Printf("Validation passed: file_path=%s, export_name=%s", args.FilePath, args.ExportName)

	return nil
}
