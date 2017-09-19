package main

import (
	"fmt"
	"os"
	tm "untitled/encoding/grpc/time"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const (
	host = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := tm.NewTimerClient(conn)

	timeZone := "UTC"
	if len(os.Args) > 1 {
		timeZone = os.Args[1]
	}

	rm, err := c.ReturnTimeNow(context.Background(), &tm.Request{TimeZone:timeZone})
	if  err != nil {
		fmt.Printf("Can`t get time: %v", err)
	}
	printErr := rm.GetError()

	if printErr == "" {
		printErr = "nil"
	}
	fmt.Printf("Time now: %s\nProccesing time: %v nsec\nError: %s\n", rm.GetNow(), rm.GetProcesingTime(), printErr)
}
