package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on UDP port: %v", err)
	}
	defer conn.Close()
	log.Println("Listening on UDP port 8080")

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("Error Reading from UDP: %v", err)
			continue
		}
		log.Printf("Received %d bytes from %s: %s", n, addr, string(buf[:n]))

		_, err = conn.WriteTo(buf[:n], addr)
		if err != nil {
			log.Printf("Error writing to %s: %v", addr, err)
		}
	}
}
