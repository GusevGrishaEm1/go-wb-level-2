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
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 30*time.Second, "timeout for connection")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Please provide host:port to connect to")
		return
	}

	address := flag.Arg(0)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			if err == io.EOF {
				conn.Close()
				return
			}
			fmt.Println("Error writing to server:", err)
		}
	}()

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Error reading from server:", err)
	}
}
