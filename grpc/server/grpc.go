package main

import (
	"fmt"
	"google.golang.org/grpc"
	"time"
	tm "untitled/encoding/grpc/time"
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
)

const (
	defaultTimeZone = "UTC"
	port = ":3000"
)

type server struct {}

func (s *server ) ReturnTimeNow(ctx context.Context, in *tm.Request) (*tm.Result, error) {
	start := time.Now()
	locStr := in.TimeZone
	loc, err := time.LoadLocation(locStr)

	if err != nil {
		fmt.Println("Can`t set location " + locStr + ", default location is \"UTC\"")
		fmt.Println(err)
		loc, err = time.LoadLocation(defaultTimeZone)
	}

	now := time.Now().In(loc)

	if err != nil {
		return &tm.Result{now.Format(time.RFC822), int64(time.Since(start)), err.Error()}, err
	}

	return &tm.Result{now.Format(time.RFC822), int64(time.Since(start)), ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tm.RegisterTimerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
