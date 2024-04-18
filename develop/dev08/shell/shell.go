package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

type command struct {
	command string
	input   string
	output  string
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		commands := strings.Split(input, "|")

		output := ""

		for _, command := range commands {
			output, err = executeCommand(command, output)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println(output)
	}
}

func executeCommand(command, output string) (string, error) {
	args := strings.Fields(command)
	var err error

	switch args[0] {
	case "cd":
		if len(args) > 1 {
			if err = os.Chdir(args[1]); err != nil {
				return "", fmt.Errorf("Error changing directory: %w", err)
			}
		} else {
			return "", fmt.Errorf("Usage: cd <directory>")
		}
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("Error getting current directory: %w", err)
		}
		output = pwd
	case "echo":
		if len(args) > 1 {
			output = strings.Join(args[1:], " ")
		} else {
			return "", fmt.Errorf("Usage: echo <text>")
		}
	case "kill":
		if len(args) > 1 {
			var pid int
			pid, err = strconv.Atoi(args[1])
			if err != nil {
				return "", fmt.Errorf("Error converting PID: %w", err)
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				return "", fmt.Errorf("Error finding process: %w", err)
			}
			err = process.Signal(syscall.SIGKILL)
			if err != nil {
				return "", fmt.Errorf("Error sending signal: %w", err)
			}
			output = fmt.Sprintf("Process with pid %d killed", pid)
		} else {
			return "", fmt.Errorf("Usage: kill <pid>")
		}
	case "ps":
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("tasklist")
		} else {
			cmd = exec.Command("ps", "aux")
		}
		outputBytes, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("error: %w", err)
		}
		output = string(outputBytes)
	case "fork":
		fmt.Print("fork")
	case "exec":
		fmt.Print("exec")
	default:
		return "", fmt.Errorf("Unknown command: %s", args[0])
	}

	return output, nil
}
