package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	TypeData = 0x01
	TypeAck  = 0x02
	Timeout  = 2 * time.Second
)

var (
	unacked      = make(map[uint32]chan bool)
	unackedMutex = sync.Mutex{}
	seqNum       uint32
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
	conn, _ := net.ListenPacket("udp", ":0")

	go func() {
		buf := make([]byte, 5)
		for {
			_, _, err := conn.ReadFrom(buf)
			if err != nil {
				continue
			}
			if buf[0] == TypeAck {
				ackNum := binary.BigEndian.Uint32(buf[1:5])
				unackedMutex.Lock()
				if ch, ok := unacked[ackNum]; ok {
					log.Printf("Received ACK for seq %d", ackNum)
					close(ch)
					delete(unacked, ackNum)
				}
				unackedMutex.Unlock()
			}
		}
	}()

	for i := 1; i <= 5; i++ {
		seqNum++
		currentSeq := seqNum
		message := fmt.Sprintf("Message %d", currentSeq)

		ackChan := make(chan bool)
		unackedMutex.Lock()
		unacked[currentSeq] = ackChan
		sendWithRetry(conn, serverAddr, seqNum, []byte(message), ackChan)
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
}

func sendWithRetry(conn net.PacketConn, addr net.Addr, seq uint32, payload []byte, ackChan chan bool) {
	packet := make([]byte, 5+len(payload))
	packet[0] = TypeData
	binary.BigEndian.PutUint32(packet[1:5], seq)
	copy(packet[5:], payload)
	for {
		log.Printf("Seding packet with seq %d", seq)
		conn.WriteTo(packet, addr)

		select {
		case <-ackChan:
			return
		case <-time.After(Timeout):
		}

	}
}
