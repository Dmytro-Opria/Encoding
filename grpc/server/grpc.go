package main

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"io"
	"net/http"
	"time"
	tm "untitled/encoding/grpc/time"
)

var (
	port = ":3005"
)

const defaultTimeZone = "UTC"

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, r.Body)

	bufStr := string(buf.Bytes()[2:])

	loc, err := time.LoadLocation(bufStr)

	if err != nil {
		fmt.Println("Can`t set location " + bufStr + ", default location is \"UTC\"")
		fmt.Println(err)
		loc, _ = time.LoadLocation(defaultTimeZone)
	}

	now := time.Now().In(loc)

	res := &tm.Result{now.Format(time.RFC822), int64(time.Since(now))}

	protoResult, err := proto.Marshal(res)
	if err != nil {
		fmt.Println("Can`t marshal", err)
	}
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(protoResult)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}
