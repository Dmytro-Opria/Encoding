package main

import (
	"net/http"
	"fmt"
	tm "untitled/encoding/grpc/time"
	"github.com/gogo/protobuf/proto"
	"bytes"
	"io"
	"os"
)

func main() {
	timeZone := "UTC"

	if len(os.Args) > 1 {
		timeZone = os.Args[1]
		fmt.Println([]byte(timeZone))
	}

	reqStr := &tm.Request{timeZone}
	protoBytes, err := proto.Marshal(reqStr)
	if err != nil {
		fmt.Println("Can`n marshal request", err)
	}
	req, err := http.NewRequest("POST", "http://localhost:3005/", bytes.NewBuffer(protoBytes))
	if  err != nil {
		fmt.Println("Can`t make post request", err)
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
	}

	newResult := &tm.Result{}

	buf := bytes.NewBuffer(nil)

	io.Copy(buf, resp.Body)

	err = proto.Unmarshal(buf.Bytes(), newResult)

	if err != nil {
		fmt.Println("Can`t unmarshal", err)
	}
	fmt.Printf("Time now: %s\nProccesing time: %v nsec\n",newResult.Now, newResult.ProcesingTime)
}
