package main

import (
	"log"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8800")
	
	if err != nil {
		log.Fatalf("Dial error: %s\n", err)
	}

	// _, err = conn.Write([]byte("Connect success\n"))
	if err != nil {
		// connection error
		log.Printf("Connection error: %s\n", err)
		conn.Close()
		return
	}

	for {
		
		inputReader := bufio.NewReader(os.Stdin)
		log.Printf("Enter a string:\n")
		inputStr, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("Input error: %s\n", err)
		}
		// log.Printf("The str is: %s\n", string(inputStr))
		conn.Write([]byte(string(inputStr)))
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
		log.Printf(string(buf[:readByte]))
	}
}