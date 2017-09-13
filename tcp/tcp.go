package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CON_TYPE = "tcp"
	CON_HOST = "localhost"
	CON_PORT = "3001"
)

func main() {
	listener, err := net.Listen(CON_TYPE, CON_HOST + ":" + CON_PORT)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on " + CON_HOST + ":" + CON_PORT)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Can`t connect", err)
			conn.Close()
			continue
		}

		fmt.Println("Connected")

		bufReader := bufio.NewReader(conn)

		fmt.Println("Start reading")

		go func(conn net.Conn) {
			defer conn.Close()

			for {
				rstring, err := bufReader.ReadString('\n')

				if err != nil {
					fmt.Println("Can`t read", err)
					break
				}

				conn.Write([]byte("Get = " + rstring))
				fmt.Print(rstring)
			}
		}(conn)

	}
}
