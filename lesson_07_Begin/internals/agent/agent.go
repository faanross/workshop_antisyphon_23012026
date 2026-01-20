package agent

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"workshop3_dev/internals/models"
)

// Agent implements the Communicator interface for HTTPS
type Agent struct {
	serverAddr string
	client     *http.Client
	// TODO: Add commandOrchestrators map to store command handlers
	// Hint: map[string]OrchestratorFunc
	commandOrchestrators map[string]OrchestratorFunc // Maps commands to their keywords
}

// NewAgent creates a new HTTPS agent
func NewAgent(serverAddr string) *Agent {
	// Create TLS config that accepts self-signed certificates
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Create HTTP client with custom TLS config
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	agent := &Agent{
		serverAddr:           serverAddr,
		client:               client,
		// TODO: Initialize commandOrchestrators map
		// Hint: make(map[string]OrchestratorFunc)
		commandOrchestrators: make(map[string]OrchestratorFunc),
	}

	// TODO: Call registerCommands to register all command handlers
	registerCommands(agent)

	return agent
}

// Send implements Communicator.Send for HTTPS
func (agent *Agent) Send(ctx context.Context) (*models.ServerResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("https://%s/", agent.serverAddr)

	// Create GET request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Send request
	resp, err := agent.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, body)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Unmarshal into ServerResponse
	var serverResp models.ServerResponse
	if err := json.Unmarshal(body, &serverResp); err != nil {
		return nil, fmt.Errorf("unmarshaling response: %w", err)
	}

	// Return the parsed response
	return &serverResp, nil
}

// TODO: Implement registerCommands to register command orchestrators
// For now, this will be empty - we'll add the shellcode orchestrator later
func registerCommands(agent *Agent) {
	// agent.commandOrchestrators["shellcode"] = (*Agent).orchestrateShellcode
	// Register other commands here in the future
}

// TODO: Implement SendResult to POST task results back to the server
// This sends the AgentTaskResult back to the server after command execution
func (agent *Agent) SendResult(resultData []byte) error {

	// TODO: Build the URL for the results endpoint
	// Hint: https://{serverAddr}/results
	targetURL := fmt.Sprintf("https://%s/results", agent.serverAddr)

	log.Printf("|RETURN RESULTS|-> Sending %d bytes of results via POST to %s", len(resultData), targetURL)

	// TODO: Create HTTP POST request with resultData as body
	// Hint: http.NewRequest(http.MethodPost, targetURL, bytes.NewReader(resultData))
	req, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewReader(resultData))
	if err != nil {
		log.Printf("|ERR SendResult| Failed to create results request: %v", err)
		return fmt.Errorf("failed to create http results request: %w", err)
	}

	// TODO: Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// TODO: Execute the request
	resp, err := agent.client.Do(req)
	if err != nil {
		log.Printf("|ERR | Results POST request failed: %v", err)
		return fmt.Errorf("http results post request failed: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("SUCCESSFULLY SENT FINAL RESULTS BACK TO SERVER.")
	return nil
}
