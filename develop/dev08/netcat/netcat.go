package netcat

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
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
