package main

import (
	"net"
	"fmt"
	"os"
)

const (
	CON_TYPE = "udp"
	CON_HOST = "localhost"
	CON_PORT = "3002"
)

func main() {
	listener, err := net.ListenPacket(CON_TYPE, CON_HOST + ":" + CON_PORT)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on " + CON_HOST + ":" + CON_PORT)

	for {

		buf  := make([]byte, 1024)

		_, addr, err := listener.ReadFrom(buf)

		if err != nil {
			fmt.Println("Can`t connect", err)
			listener.Close()
			continue
		}

		fmt.Print(string(buf))
		listener.WriteTo(append([]byte("Get = "), buf...), addr)

	}
}
