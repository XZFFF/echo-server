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

	_, err = conn.Write([]byte("Connect success\n"))
	if err != nil {
		// connection error
		log.Printf("Connection error: %s\n", err)
		conn.Close()
		return
	} else {
		log.Printf("Connect success\n")
	}

	for {
		log.Printf("Enter a string:\n")
		inputReader := bufio.NewReader(os.Stdin)
		inputStr, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("Input error: %s\n", err)
		}
		// log.Printf("The str is: %s\n", string(inputStr))
		conn.Write([]byte(string(inputStr)))
	}
}

