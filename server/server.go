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


	for {
		conn, err := ln.Accept()
		if err != nil {
			// accept error
			log.Printf("Accept error: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) { 
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
		log.Printf("My data:%s\n", string(buf[:readByte]))
		conn.Write([]byte("Echo data:\n"+string(buf[:readByte])))
	}
}
