package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		Control: func(network string, address string, c syscall.RawConn) error {
			return nil
		},
	}

	conn, err := dialer.DialContext(context.Background(), "tcp", "example.com:80")
	if err != nil {
		return
	}

	defer conn.Close()
	log.Println(" Connected to example.com with KeepAlive enabled. ")
}
