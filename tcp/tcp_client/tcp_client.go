package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)


func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:3001")
	reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("Text to send: ")
			text, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Can`t read string", err)
			}

			fmt.Fprintf(conn, text)

			buf := make([]byte, len([]byte(text))+6)
			conn.Read(buf)

			if err != nil {
				fmt.Println("Can`t read line", err)
			}
			fmt.Print("Message from server: " + string(buf))
		}
}
/*
func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:3001")
	reader := bufio.NewReader(os.Stdin)

	proto := textproto.NewReader(bufio.NewReader(conn))
	for {
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Can`t read string", err)
		}

		fmt.Fprintf(conn, text)
		line, err :=  proto.ReadLine()

		if err != nil {
			fmt.Println("Can`t read line", err)
		}
		fmt.Print("Message from server: " + line  +"\n")
	}
}*/

/*

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:3001")
	reader := bufio.NewReader(os.Stdin)
	message, _, _ := bufio.NewReader(conn).ReadLine()

	for {
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text)
		if len(message) < 7 {
			time.Sleep(time.Second/50)
			continue
		}
		fmt.Print("Message from server: " + string(message)+"\n")
	}

}
*/


