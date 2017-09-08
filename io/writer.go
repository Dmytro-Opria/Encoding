package main

import (
	"fmt"
	"io"
	"os"
)

type Writer struct {
}

type FileWrite struct {
	filename string
}

func main() {
	path := "test.txt"
	w := Writer{}
	fw := FileWrite{filename:"file.txt"}
	//w.Write([]byte("Byte array"))
	//fmt.Fprintln(w,"ABCDEFGEREDFDF")
	//io.Copy(w, strings.NewReader("ABCDEFGEREDFDF"))
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Can not open file", path, err)
	}
	defer file.Close()
	mwriter := io.MultiWriter(w, fw)
	io.Copy(mwriter, file)
}

func(w Writer) Write(p []byte) (n int, err error){
	fmt.Println(p)

	return len(p), nil
}

func(w FileWrite) Write(p []byte) (n int, err error){
	file, err := os.Create(w.filename)
	if err != nil {
		fmt.Println("Can not open file", w.filename, err)
	}
	defer file.Close()

	file.Write(p)

	return len(p), nil
}
