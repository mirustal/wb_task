package client

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.String("timeout", "10s", "connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=duration] <host> <port>")
		return
	}
	host, port := args[0], args[1]

	// Парсинг таймаута
	duration, err := time.ParseDuration(*timeout)
	if err != nil {
		fmt.Println("Invalid timeout format:", err)
		return
	}

	address := net.JoinHostPort(host, port)
	fmt.Printf("Connecting to %s...\n", address)

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Подключение к серверу
	conn, err := net.DialTimeout("tcp", address, duration)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected.")

	go readFromServer(conn)
	writeToServer(ctx, conn)
}

func readFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("Server: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Read error: %s\n", err)
	}
	fmt.Println("Server connection closed.")
}

func writeToServer(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Connection closed due to timeout.")
			return
		default:
			if !scanner.Scan() {
				fmt.Println("Exiting...")
				return
			}
			text := scanner.Text()
			_, err := conn.Write([]byte(text + "\n"))
			if err != nil {
				fmt.Printf("Write error: %s\n", err)
				return
			}
		}
	}
}
