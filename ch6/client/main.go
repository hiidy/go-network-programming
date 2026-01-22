package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to resolve server address: %v", err)
	}

	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return
	}

	defer conn.Close()
	log.Printf("Client is listening on %s", conn.LocalAddr())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text to send to the server (type 'exit' to quit):")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		_, err := conn.WriteTo([]byte(text), serverAddr)
		if err != nil {
			log.Printf("Failed to write to server: %v", err)
			continue
		}

		buf := make([]byte, 1024)
		n, _, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("Failed to read from server: %v", err)
			continue
		}
		fmt.Printf("Echo from server: %s\n", string(buf[:n]))
	}
}
