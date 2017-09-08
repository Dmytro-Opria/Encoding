package main

import (
	_"io"
	"fmt"
)

type Writer struct {
	data []byte
}

func main() {
	w := Writer{data: []byte("")}
	w.Write([]byte("Byte array"))
}

func(w Writer) Write(p []byte) (n int, err error){
	fmt.Println(p)
	return 0, nil
}
