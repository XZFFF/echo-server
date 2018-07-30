package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8800")
	
	if err != nil {
		log.Fatalf("Dial error: %s\n", err)
	}

	_, err = conn.Write([]byte("Try connect...\n"))

	if err != nil {
		// connection error
		log.Printf("Connection error: %s\n", err)
		conn.Close()
		return
	}
	
}