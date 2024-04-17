package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}

	var timeout time.Duration
	if strings.Contains(args[0], "--timeout") {
		timeout, _ = time.ParseDuration(strings.Split(args[0], "--timeout=")[1])
	}

	host := args[1]
	port := args[2]

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Print("connected: ", conn.RemoteAddr().String())

	go func() {
		fmt.Fprint(os.Stdout, "sdf")
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("Failed to copy from connection:", err)
			os.Exit(1)
		}
	}()

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Failed to copy to connection:", err)
			os.Exit(1)
		}
	}()

	<-make(chan struct{})
}
