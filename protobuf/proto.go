package main

import (
	"log"
	"github.com/golang/protobuf/proto"
	"fmt"
	"untitled/encoding/protobuf/users"
	"os"
	"io"
	"bytes"
)

var user = &users.User {
	Name: "John Doe",
	Age:  20,
	Documents:  []string{"doc_A","doc_B","doc_C"},
}
var path = "proto"

func main() {
	data, _ := encode()

	writeData(data)

	data, _ = readData()

	newUser, _ := decode(data)
	fmt.Println(newUser)
}

func encode()(data []byte,err error){
	data, err = proto.Marshal(user)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return
}

func decode(data []byte) (newUser *users.User, err error){
	newUser = &users.User{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	return
}

func writeData (data []byte)(err error){
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Can`t create file", err)
	}

	file.Write(data)

	return
}

func readData()(data []byte, err error){
	buf := bytes.NewBuffer(nil)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Can`t open file", err)
	}
	defer file.Close()

	io.Copy(buf, file)

	data = buf.Bytes()

	return
}