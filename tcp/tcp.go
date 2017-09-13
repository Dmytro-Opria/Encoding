package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

const (
	CON_TYPE = "tcp"
	CON_HOST = "localhost"
	CON_PORT = "3001"
)

func main() {
	l, err := net.Listen(CON_TYPE, CON_HOST + ":" + CON_PORT)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + CON_HOST + ":" + CON_PORT)

	conn, err := l.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message Received:", string(message))

		conn.Write([]byte(message + "\n"))
	}
}