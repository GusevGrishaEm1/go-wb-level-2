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
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Please provide host:port to connect to")
		return
	}

	address := flag.Arg(0)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		<-time.After(*timeout)
		fmt.Println("Программа завершается по таймауту")
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Завершение по сигналу")
		conn.Close()
		os.Exit(1)
	}()

	if conn != nil {
		defer conn.Close()

		go func() {
			defer wg.Done()
			reader := bufio.NewReader(conn)
			for {
				data, err := reader.ReadBytes('\n')
				if err != nil {
					fmt.Println("Error reading from server:", err)
					return
				}
				fmt.Print(string(data))
			}
		}()
		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				input := scanner.Text()
				_, err := conn.Write([]byte(input + "\n"))
				if err != nil {
					fmt.Println("Error writing to server:", err)
					break
				}
			}
		}()

	}

	wg.Wait()
}
