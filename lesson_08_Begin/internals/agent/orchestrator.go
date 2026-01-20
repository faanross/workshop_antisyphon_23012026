package agent

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"workshop3_dev/internals/models"
)

// TODO: Implement orchestrateShellcode to handle the "shellcode" command
// This function receives the job from the server, extracts shellcode args,
// decodes the base64 shellcode, and calls the actual shellcode execution
func (agent *Agent) orchestrateShellcode(job *models.ServerResponse) models.AgentTaskResult {

	// TODO: Create a variable to hold the shellcode arguments
	var shellcodeArgs models.ShellcodeArgsAgent

	// TODO: Unmarshal job.Arguments into shellcodeArgs
	// Hint: json.Unmarshal(job.Arguments, &shellcodeArgs)
	if err := json.Unmarshal(job.Arguments, &shellcodeArgs); err != nil {
		errMsg := fmt.Sprintf("Failed to unmarshal ShellcodeArgs for Task ID %s: %v. ", job.JobID, err)
		log.Printf("|ERR SHELLCODE ORCHESTRATOR| %s", errMsg)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("failed to unmarshal ShellcodeArgs"),
		}
	}
	log.Printf("|SHELLCODE ORCHESTRATOR| Task ID: %s. Executing Shellcode, Export Function: %s, ShellcodeLen(b64)=%d\n",
		job.JobID, shellcodeArgs.ExportName, len(shellcodeArgs.ShellcodeBase64))

	// TODO: Validate that ShellcodeBase64 is not empty
	if shellcodeArgs.ShellcodeBase64 == "" {
		log.Printf("|ERR SHELLCODE ORCHESTRATOR| Task ID %s: ShellcodeBase64 is empty.", job.JobID)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("ShellcodeBase64 cannot be empty"),
		}
	}

	// TODO: Validate that ExportName is not empty
	if shellcodeArgs.ExportName == "" {
		log.Printf("|ERR SHELLCODE ORCHESTRATOR| Task ID %s: ExportName is empty.", job.JobID)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("ExportName must be specified for DLL execution"),
		}
	}

	// TODO: Decode the base64 shellcode
	// Hint: base64.StdEncoding.DecodeString(shellcodeArgs.ShellcodeBase64)
	rawShellcode, err := base64.StdEncoding.DecodeString(shellcodeArgs.ShellcodeBase64)
	if err != nil {
		log.Printf("|ERR SHELLCODE ORCHESTRATOR| Task ID %s: Failed to decode ShellcodeBase64: %v", job.JobID, err)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("Failed to decode shellcode"),
		}
	}

	// TODO: Call the "doer" function (will be implemented in lesson 09)
	// For now, this is a placeholder - the actual shellcode package will be added later
	// commandShellcode := shellcode.New()
	// shellcodeResult, err := commandShellcode.DoShellcode(rawShellcode, shellcodeArgs.ExportName)

	// Placeholder result for now
	log.Printf("|SHELLCODE ORCHESTRATOR| Decoded shellcode: %d bytes", len(rawShellcode))

	finalResult := models.AgentTaskResult{
		JobID: job.JobID,
	}

	// The actual implementation will set Success and CommandResult based on shellcode execution
	// For now, just return a placeholder
	outputJSON, _ := json.Marshal("Shellcode orchestrator placeholder - doer not yet implemented")
	finalResult.CommandResult = outputJSON

	return finalResult
}
