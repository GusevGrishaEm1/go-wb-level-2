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

		args := strings.Fields(input)

		switch args[0] {
		case "cd":
			if len(args) > 1 {
				if err := os.Chdir(args[1]); err != nil {
					fmt.Println("Error changing directory:", err)
				}
			} else {
				fmt.Println("Usage: cd <directory>")
			}
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
			} else {
				fmt.Println(pwd)
			}
		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			} else {
				fmt.Println("Usage: echo <text>")
			}
		case "kill":
			if len(args) > 1 {
				var pid int
				pid, err = strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Error converting PID:", err)
					return
				}
				process, err := os.FindProcess(pid)
				if err != nil {
					fmt.Println("Error finding process:", err)
					return
				}
				err = process.Signal(syscall.SIGKILL)
				if err != nil {
					fmt.Println("Error send signal:", err)
					return
				}
				fmt.Println("Process with pid ", pid, "killed")
			} else {
				fmt.Println("Usage: kill <pid>")
			}
		case "ps":
			var cmd *exec.Cmd
			if runtime.GOOS == "windows" {
				cmd = exec.Command("tasklist")
			} else {
				cmd = exec.Command("ps", "aux")
			}
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("error: ", err)
			}
			fmt.Println(string(output))
		default:
			pid, err := syscall.Fork()
			if err != nil {
				fmt.Println("Error:", err)
			}
			if pid == 0 {
				err = syscall.Exec("/bin/ls", []string{"ls"}, nil)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
			fmt.Println("PID:", pid)
		}
	}
}
