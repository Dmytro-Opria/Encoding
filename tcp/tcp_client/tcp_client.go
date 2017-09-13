package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"time"
)


func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:3001")
	reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')

			fmt.Fprintf(conn, text)

			message, _, _ := bufio.NewReader(conn).ReadLine()
			if len(message) < 7 {
				time.Sleep(time.Second/50)
				continue
			}
			fmt.Print("Message from server: " + string(message)+"\n")
		}

}


