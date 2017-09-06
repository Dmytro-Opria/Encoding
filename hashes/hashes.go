package main

import (
	"io"
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"encoding/base64"
	"crypto/sha1"
)

func main(){
	testStr := "abcdefg"

	fmt.Println(toMd5(testStr))
	fmt.Println(toSha1(testStr))
}

func toMd5(input string)(bytes []byte, hexStr string, base64Str string){
	h := md5.New()
	io.WriteString(h,input)

	bytes = h.Sum(nil)

	hexStr = hex.EncodeToString(bytes)

	base64Str = base64.StdEncoding.EncodeToString(bytes)

	return
}

func toSha1(input string)(bytes []byte, hexStr string, base64Str string){
	h := sha1.New()
	io.WriteString(h,input)

	bytes = h.Sum(nil)

	hexStr = hex.EncodeToString(bytes)

	base64Str = base64.StdEncoding.EncodeToString(bytes)

	return
}

