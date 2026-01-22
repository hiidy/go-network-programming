package main

import (
	"encoding/binary"
	"log"
	"net"
)

const (
	TypeData = 0x01
	TypeAck  = 0x02
)

func main() {
	conn, _ := net.ListenPacket("udp", ":8080")
	defer conn.Close()
	log.Println("Reliable UDP Server listening on : 8080")

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}

		if n > 5 && buf[0] == TypeData {
			seqNum := binary.BigEndian.Uint32(buf[1:5])
			log.Printf("Received data packet with seq %d from %s", seqNum, remoteAddr)

			ackPacket := make([]byte, 5)
			ackPacket[0] = TypeAck
			binary.BigEndian.PutUint32(ackPacket[1:5], seqNum)
			conn.WriteTo(ackPacket, remoteAddr)
		}
	}
}
