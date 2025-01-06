package common

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
)

// ExecuteOption 定义执行选项的函数类型
type ExecuteOption func(*executeOptions)

// executeOptions 内部选项结构体
type executeOptions struct {
	printCommand bool   // 是否打印执行的命令
	verbose      bool   // 是否输出详细信息
	outputFormat string // 输出格式：text/json
	timeout      int    // 执行超时时间(秒)
}

// 定义选项设置函数
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

// ExecuteResult 定义执行结果
type ExecuteResult struct {
	Command  string   `json:"command"`
	Args     []string `json:"args"`
	Output   string   `json:"output"`
	Error    string   `json:"error,omitempty"`
	ExitCode int      `json:"exit_code"`
}

// 命令执行器接口
type Executor interface {
	Execute(command string, args []string, opts ...ExecuteOption) (*ExecuteResult, error)
}

// CephExecutor 实现
type CephExecutor struct{}

func NewCephExecutor() *CephExecutor {
	return &CephExecutor{}
}

func (e *CephExecutor) Execute(command string, args []string, opts ...ExecuteOption) (*ExecuteResult, error) {
	// 设置默认选项
	options := &executeOptions{
		outputFormat: "text",
		timeout:      30,
	}

	// 应用所有选项
	for _, opt := range opts {
		opt(options)
	}

	result := &ExecuteResult{
		Command: command,
		Args:    args,
	}

	// 打印命令
	if options.printCommand {
		fmt.Printf("Executing: %s %v\n", command, args)
	}

	// 实现具体的命令执行逻辑
	cephCmd := cmd.NewCmd(command, args...)

	statusChan := cephCmd.Start() // non-blocking

	ticker := time.NewTicker(1 * time.Second)
	// 创建一个用于超时的 channel
	timeoutChan := time.After(time.Duration(options.timeout) * time.Second)

	// Print last line of stdout every 1s
	go func() {
		for range ticker.C {
			status := cephCmd.Status()
			n := len(status.Stdout)
			fmt.Println(status.Stdout[n-1])
		}
	}()

	// 等待命令完成或超时
	select {
	case finalStatus := <-statusChan:
		// 命令正常完成
		ticker.Stop()
		result.Output = strings.Join(finalStatus.Stdout, "\n")
		result.Error = strings.Join(finalStatus.Stderr, "\n")
		result.ExitCode = finalStatus.Exit
	case <-timeoutChan:
		// 命令超时
		ticker.Stop()
		cephCmd.Stop()
		result.Error = fmt.Sprintf("command execution timed out after %d seconds", options.timeout)
		result.ExitCode = -1
		return result, fmt.Errorf("command execution timed out after %d seconds", options.timeout)
	}

	// 处理输出格式
	if options.outputFormat == "json" {
		jsonBytes, err := json.Marshal(result)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal result to JSON: %v", err)
		}
		fmt.Println(string(jsonBytes))
	}

	// 详细输出
	if options.verbose {
		fmt.Printf("Command result: %+v\n", result.Output)
		fmt.Printf("Command error: %+v\n", result.Error)
		fmt.Printf("Command exit code: %+v\n", result.ExitCode)
	}

	return result, nil
}
