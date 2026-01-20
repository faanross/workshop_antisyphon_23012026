package agent

import (
	"encoding/json"
	"errors"
	"log"
	"workshop3_dev/internals/models"
)

// TODO: Define OrchestratorFunc as a function type for command handlers
// Each command (shellcode, etc.) will have its own orchestrator function
// Hint: func(agent *Agent, job *models.ServerResponse) models.AgentTaskResult
type OrchestratorFunc func(agent *Agent, job *models.ServerResponse) models.AgentTaskResult

// TODO: Implement ExecuteTask method to dispatch commands to their orchestrators
// This is called when the agent receives a job from the server
func (agent *Agent) ExecuteTask(job *models.ServerResponse) {
	log.Printf("AGENT IS NOW PROCESSING COMMAND %s with ID %s", job.Command, job.JobID)

	var result models.AgentTaskResult

	// TODO: Look up the orchestrator for this command
	// Hint: orchestrator, found := agent.commandOrchestrators[job.Command]
	orchestrator, found := agent.commandOrchestrators[job.Command]

	if found {
		// TODO: Call the orchestrator to execute the command
		result = orchestrator(agent, job)
	} else {
		log.Printf("|WARN AGENT TASK| Received unknown command: '%s' (ID: %s)", job.Command, job.JobID)
		result = models.AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   errors.New("command not found"),
		}
	}

	// TODO: Marshal the result to JSON
	resultBytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("|ERR AGENT TASK| Failed to marshal result for Task ID %s: %v", job.JobID, err)
		return // Cannot send result if marshalling fails
	}

	// TODO: Send the result back to the server using SendResult
	log.Printf("|AGENT TASK|-> Sending result for Task ID %s (%d bytes)...", job.JobID, len(resultBytes))
	err = agent.SendResult(resultBytes)
	if err != nil {
		log.Printf("|ERR AGENT TASK| Failed to send result for Task ID %s: %v", job.JobID, err)
	}

	log.Printf("|AGENT TASK|-> Successfully sent result for Task ID %s.", job.JobID)

}
