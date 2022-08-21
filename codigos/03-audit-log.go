package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	count := 0
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		count += 1
		fmt.Printf("number of activr %d\n", count)
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {

	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().String()+"\n")
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func sendAuditLog(chClient <-chan string) {
	for {
		select {
		case msg := <-chClient:
			fmt.Print(msg)
		default:

		}
	}
}
