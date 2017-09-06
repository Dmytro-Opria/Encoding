package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"os"
	"bufio"
)

func main() {
	//testStr := "abcdefg"
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(toMd5(strings.NewReader(text)))
	//fmt.Fprintln(os.Stdout, "test")
	//fmt.Println(toSha1(strings.NewReader(testStr)))
}

func toMd5(r io.Reader) (bytes []byte, hexStr string, base64Str string) {

	h := md5.New()

	io.Copy(h, r)

	bytes = h.Sum(nil)

	hexStr = hex.EncodeToString(bytes)

	base64Str = base64.StdEncoding.EncodeToString(bytes)

	return
}

func toSha1(r io.Reader) (bytes []byte, hexStr string, base64Str string) {

	h := sha1.New()

	io.Copy(h, r)

	bytes = h.Sum(nil)

	hexStr = hex.EncodeToString(bytes)

	base64Str = base64.StdEncoding.EncodeToString(bytes)

	return
}
