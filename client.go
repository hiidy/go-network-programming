package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	log.Println("connected to server")

	scanner := bufio.NewScanner(os.Stdin)
	serverReader := bufio.NewReader(conn)

	fmt.Println("Enter text to send to the server (type 'exit' to quit):")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			log.Printf("Failed to write to server: %v", err)
			return
		}

		response, err := serverReader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read from server: %v", err)
			return
		}

		fmt.Printf("Echo from server: %s", response)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from stdin: %v", err)
	}

	log.Println("Connection closed.")
}
