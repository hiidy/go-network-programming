package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Dialing...")
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	defer conn.Close()
	log.Printf("Connection established to %s", conn.RemoteAddr())
	time.Sleep(10 * time.Second)
}
