package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Создание канала для получения сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Print("$ ")
	var command string
	fmt.Scanln(&command)

	// Разделение команды на аргументы
	args := splitCommand(command)

	// Обработка встроенных команд
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
		fmt.Println(args[1])
	case "kill":
		// TODO: Реализовать функцию kill
	case "ps":
		// TODO: Реализовать функцию ps
	default:
		// Выполнение команды с помощью fork/exec
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
	}

	protocol, address := parseFlagsForNetcat()
	err := netcat(protocol, address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
}

func parseFlagsForNetcat() (string, string) {
	protocol := *flag.String("protocol", "tcp", "Connection protocol (tcp or udp)")
	address := *flag.String("address", "localhost:8080", "Address to connect to")
	flag.Parse()
	return protocol, address
}

func netcat(protocol, address string) error {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("Error sending data:", err)
			os.Exit(1)
		}
	}
	if scanner.Err() != nil {
		fmt.Println("Error reading from stdin:", scanner.Err())
		os.Exit(1)
	}
	return nil
}
