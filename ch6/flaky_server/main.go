package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const dropRate = 0.3

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		return
	}

	defer conn.Close()
	log.Println("Flaky UDP Server listening on :8080")

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}

		if rand.Float32() < dropRate {
			log.Printf("dropped packet from %s", remoteAddr)
			continue
		}

		log.Printf("Received %d bytes from %s: %s", n, remoteAddr, string(buf[:n]))

		_, err = conn.WriteTo(buf[:n], remoteAddr)
		if err != nil {
			log.Printf("Error Writing to %s: %v", remoteAddr, err)
		}
	}
}
