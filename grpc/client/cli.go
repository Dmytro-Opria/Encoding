package main

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"io"
	"net/http"
	"os"
	tm "untitled/encoding/grpc/time"
)

func main() {
	timeZone := "UTC"

	if len(os.Args) > 1 {
		timeZone = os.Args[1]
	}

	reqStr := &tm.Request{timeZone}
	protoBytes, err := proto.Marshal(reqStr)
	if err != nil {
		fmt.Println("Can`t marshal request", err)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:3005/", bytes.NewBuffer(protoBytes))
	if err != nil {
		fmt.Println("Can`t make post request", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-protobuf")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Can`t get response from server", err)
		return
	}

	newResult := &tm.Result{}

	buf := bytes.NewBuffer(nil)

	io.Copy(buf, resp.Body)

	err = proto.Unmarshal(buf.Bytes(), newResult)

	if err != nil {
		fmt.Println("Can`t unmarshal", err)
		return 
	}
	fmt.Printf("Time now: %s\nProccesing time: %v nsec\n", newResult.Now, newResult.ProcesingTime)
}
