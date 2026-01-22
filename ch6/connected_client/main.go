package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Printf("Client connected to %s from %s", conn.RemoteAddr(), conn.LocalAddr())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text to send to the server (type 'exit' to quit):")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Printf("Failed to write to server: %v", err)
			continue
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read from server: %v", err)
			break
		}
		fmt.Printf("Echo from server: %s\n", string(buf[:n]))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from stdin: %v", err)
	}

	log.Println("Connection closed.")
}
