package main

import (
	"io"
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"encoding/base64"
	"crypto/sha1"
	"strings"
)

func main(){
	testStr := "abcdefg"

	fmt.Println(toMd5(strings.NewReader(testStr)))

	fmt.Println(toSha1(strings.NewReader(testStr)))
}

func toMd5(r io.Reader)(bytes []byte, hexStr string, base64Str string){
	input := make([]byte, 1000)

	r.Read(input)

	h := md5.New()
	io.WriteString(h,string(input))

	bytes = h.Sum(nil)

	hexStr = hex.EncodeToString(bytes)

	base64Str = base64.StdEncoding.EncodeToString(bytes)

	return
}

func toSha1(r io.Reader)(bytes []byte, hexStr string, base64Str string){
	input := make([]byte, 1000)

	r.Read(input)

	h := sha1.New()
	io.WriteString(h,string(input))

	bytes = h.Sum(nil)

	hexStr = hex.EncodeToString(bytes)

	base64Str = base64.StdEncoding.EncodeToString(bytes)

	return
}

