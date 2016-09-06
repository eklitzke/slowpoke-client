package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func run(conn net.Conn, messages chan string) {
	send := make([]byte, 1)
	send[0] = '.'
	buf := make([]byte, 4096)
	for {
		nbytes, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("Got read error: %v\n", err)
		}
		line := string(buf[:nbytes])
		var timeout float64
		var score, maxScore int
		fmt.Sscanf(line, "%f %d %d\n", &timeout, &score, &maxScore)
		time.Sleep(time.Duration(timeout*1000000) * time.Microsecond)
		nbytes, err = conn.Write(send)
		if err != nil {
			log.Fatalf("Got write error: %v\n", err)
		}
	}
	messages <- "all done"
}

func main() {
	messages := make(chan string)
	for i := 0; i < 1000; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:9000")
		if err != nil {
			log.Fatalf("Got dial error: %v\n", err)
		}
		go run(conn, messages)
	}
	msg := <-messages
	fmt.Println(msg)
}
