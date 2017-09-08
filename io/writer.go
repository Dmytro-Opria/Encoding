package main

import (
	_"io"
	"fmt"
)

type Writer struct {
}

func main() {
	w := Writer{}
	w.Write([]byte("Byte array"))
}

func(w Writer) Write(p []byte) (n int, err error){
	fmt.Println(p)
	return len(p), nil
}
