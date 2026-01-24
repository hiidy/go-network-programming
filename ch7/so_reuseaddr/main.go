package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	lc := net.ListenConfig{
		Control: func(network string, address string, c syscall.RawConn) error {
			var opErr error
			err := c.Control(func(fd uintptr) {
				log.Printf("Setting SO_RESUSEADDR on fd %d", fd)
				opErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
			})
			if err != nil {
				return err
			}
			return opErr
		},
	}

	listener, err := lc.Listen(context.Background(), "tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}

	defer listener.Close()
	log.Println("Listening on :8080 with SO_REUSEADDR")

	// Simulate doing some work and then exiting, to test quick restart
	time.Sleep(10 * time.Second)
	log.Println("Server shutting down.")
}
