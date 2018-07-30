package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	log.Printf("Listening on TCP prot: 8800")

	ln, err := net.Listen("tcp", ":8800")
	// listen error
	if err != nil {
		log.Fatalf("Listen error: %s\n", err)
	}

	accepted := 0 
	for {
		conn, err := ln.Accept()
		if err != nil {
			// accept error
			log.Printf("Accept error: %s\n", err)
			continue
		}
		accepted++
		go handleConnection(conn, accepted)
	}
}

func handleConnection(conn net.Conn, accepted int) { 
	bufread := bufio.NewReader(conn)
	buf := make([]byte, 1024)

	for {
		readByte, err := bufread.Read(buf)
		// connection error
		if err != nil {
			log.Printf("Connection error: %s\n", err)
			conn.Close()
			return
		}
		log.Printf("Client%d data:%s", accepted, string(buf[:readByte]))
		conn.Write([]byte("Server echo data:" + string(buf[:readByte])))
	}
}
