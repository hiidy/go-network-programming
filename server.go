package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	defer listener.Close()
	log.Println("listening on 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept conn: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Client connected: %s", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("Client disconnected: %s", conn.RemoteAddr())
			} else {
				log.Printf("Error reading from client: %v", err)
			}
			return
		}

		log.Printf("Received: %s", string(message))
		conn.Write([]byte(message))
	}
}
