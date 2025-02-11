package common

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
)

// ExecuteOption defines the function type for execution options
type ExecuteOption func(*executeOptions)

// executeOptions internal options struct
type executeOptions struct {
	printCommand bool   // Whether to print the executed command
	verbose      bool   // Whether to output detailed information
	outputFormat string // Output format: text/json
	timeout      int    // Execution timeout (seconds)
}

// Define option setting functions
func WithPrintCommand() ExecuteOption {
	return func(o *executeOptions) {
		o.printCommand = true
	}
}

func WithVerbose() ExecuteOption {
	return func(o *executeOptions) {
		o.verbose = true
	}
}

func WithOutputFormat(format string) ExecuteOption {
	return func(o *executeOptions) {
		o.outputFormat = format
	}
}

func WithTimeout(seconds int) ExecuteOption {
	return func(o *executeOptions) {
		o.timeout = seconds
	}
}

// ExecuteResult defines the execution result
type ExecuteResult struct {
	Command  string   `json:"command"`
	Args     []string `json:"args"`
	Output   string   `json:"output"`
	Error    string   `json:"error,omitempty"`
	ExitCode int      `json:"exit_code"`
}

// Command executor interface
type Executor interface {
	Execute(command string, args []string, opts ...ExecuteOption) (*ExecuteResult, error)
}

// CephExecutor implementation
type CephExecutor struct{}

func NewCephExecutor() *CephExecutor {
	return &CephExecutor{}
}

func (e *CephExecutor) Execute(command string, args []string, opts ...ExecuteOption) (*ExecuteResult, error) {
	// Set default options
	options := &executeOptions{
		outputFormat: "text",
		timeout:      30,
	}

	// Apply all options
	for _, opt := range opts {
		opt(options)
	}

	result := &ExecuteResult{
		Command: command,
		Args:    args,
	}

	// Print command
	if options.printCommand {
		fmt.Printf("Executing: %s %v\n", command, args)
	}

	// Implement the specific command execution logic
	cephCmd := cmd.NewCmd(command, args...)

	statusChan := cephCmd.Start() // non-blocking

	ticker := time.NewTicker(1 * time.Second)
	// Create a channel for timeout
	timeoutChan := time.After(time.Duration(options.timeout) * time.Second)

	// Print last line of stdout every 1s
	go func() {
		for range ticker.C {
			status := cephCmd.Status()
			n := len(status.Stdout)
			fmt.Println(status.Stdout[n-1])
		}
	}()

	// Wait for the command to complete or timeout
	select {
	case finalStatus := <-statusChan:
		// Command completed normally
		ticker.Stop()
		result.Output = strings.Join(finalStatus.Stdout, "\n")
		result.Error = strings.Join(finalStatus.Stderr, "\n")
		result.ExitCode = finalStatus.Exit
	case <-timeoutChan:
		// Command timed out
		ticker.Stop()
		cephCmd.Stop()
		result.Error = fmt.Sprintf("command execution timed out after %d seconds", options.timeout)
		result.ExitCode = -1
		return result, fmt.Errorf("command execution timed out after %d seconds", options.timeout)
	}

	// Handle output format
	if options.outputFormat == "json" {
		jsonBytes, err := json.Marshal(result)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal result to JSON: %v", err)
		}
		fmt.Println(string(jsonBytes))
	}

	// Detailed output
	if options.verbose {
		fmt.Printf("Command result: %+v\n", result.Output)
		fmt.Printf("Command error: %+v\n", result.Error)
		fmt.Printf("Command exit code: %+v\n", result.ExitCode)
	}

	return result, nil
}
