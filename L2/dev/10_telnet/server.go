package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}
	defer listener.Close()
	log.Println("Listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received:", text)

		response := strings.ToUpper(text)
		conn.Write([]byte(response + "\n"))
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v", err)
	}
	fmt.Println("Connection closed")
}
