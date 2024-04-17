package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from the client
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		data := scanner.Text()
		fmt.Println("Received:", data)

		// Echo back the received data to the client
		_, err := conn.Write([]byte(data + "\n"))
		if err != nil {
			fmt.Println("Failed to write to connection:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}
}

func main() {
	// Specify the port number on which the server will listen
	port := "8080"

	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on port", port)

	// Accept incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
