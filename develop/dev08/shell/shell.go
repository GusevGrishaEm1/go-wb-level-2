package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		executeCommand(strings.Join(os.Args[1:], " "))
		return
	}
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		input = strings.TrimSpace(input)

		if strings.Contains(input, "|") {
			parts := strings.Split(input, "|")
			var commands [][]string
			for _, part := range parts {
				commands = append(commands, strings.Fields(strings.TrimSpace("go run shell.go "+part)))
			}
			executePipeline(commands)
		} else {
			executeCommand(input)
		}
	}
}

func executeCommand(command string) {
	args := strings.Fields(command)
	var err error

	switch args[0] {
	case "cd":
		if len(args) > 1 {
			if err = os.Chdir(args[1]); err != nil {
				fmt.Errorf("Error changing directory: %w", err)
			}
		} else {
			fmt.Errorf("Usage: cd <path>")
		}
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Errorf("Error getting current directory: %w", err)
		}
		fmt.Println(pwd)
	case "echo":
		if len(args) > 1 {
			fmt.Println(strings.Join(args[1:], " "))
		} else {
			fmt.Errorf("Usage: echo <text>")
		}
	case "kill":
		if len(args) > 1 {
			kill(args[1])
		} else {
			fmt.Errorf("Usage: kill <pid>")
		}
	case "exit":
		os.Exit(0)
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("ps:", err)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Println("ps:", err)
		}
	case "fork_exec":
		cmd := exec.Command("go", "run", "shell.go", strings.Join(args[1:], " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		fmt.Errorf("Unknown command: %s", args[0])
	}
}

func executePipeline(commands [][]string) {
	if len(commands) == 0 {
		return
	}

	cmds := make([]*exec.Cmd, len(commands))
	for i, cmdArgs := range commands {
		cmds[i] = exec.Command(cmdArgs[0], cmdArgs[1:]...)
	}

	for i := 0; i < len(cmds)-1; i++ {
		pipe, err := cmds[i].StdoutPipe()
		if err != nil {
			fmt.Println("Error creating pipe:", err)
			return
		}
		cmds[i+1].Stdin = pipe
	}

	cmds[len(cmds)-1].Stdout = os.Stdout

	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error starting command:", err)
			return
		}
	}

	for _, cmd := range cmds {
		err := cmd.Wait()
		if err != nil {
			fmt.Println("Error waiting for command:", err)
		}
	}
}

func kill(pid string) {
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("kill:", err)
		return
	}

	process, err := os.FindProcess(pidInt)
	if err != nil {
		fmt.Println("kill:", err)
		return
	}

	err = process.Kill()
	if err != nil {
		fmt.Println("kill:", err)
	}
}
