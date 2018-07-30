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

	// 循环读入信息并传输
	for {
		// 接收命令行的输入	
		log.Printf("Enter a string:\n")
		inputReader := bufio.NewReader(os.Stdin)
		inputStr, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("Input error: %s\n", err)
		}
		// log.Printf("The str is: %s\n", string(inputStr))
		// 传输信息到server
		conn.Write([]byte(string(inputStr)))

		// 获取server传输过来的消息
		bufread := bufio.NewReader(conn)
		buf := make([]byte, 1024)
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
