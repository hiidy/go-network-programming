package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("listent failed: %v", err)
	}
	defer listener.Close()
	log.Println("Listening on 127.0.0.1:8080")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Accept failed: %v", err)
	}

	defer conn.Close()

	log.Printf("connection accepted from %s", conn.RemoteAddr())

	time.Sleep(10 * time.Second)
	log.Println("Closing connection.")
}
