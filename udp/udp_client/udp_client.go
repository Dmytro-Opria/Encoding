package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:3002")
	if err != nil {
		fmt.Println("Can't connect to server", err)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("Can't connect to local", err)
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		fmt.Println("Can't make connect", err)
	}

	defer Conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Can`t read string", err)
		}

		fmt.Fprintf(Conn, text)

		buf := make([]byte, len([]byte(text))+6)
		Conn.Read(buf)

		if err != nil {
			fmt.Println("Can`t read line", err)
		}
		fmt.Print("Message from server: " + string(buf))
	}
}