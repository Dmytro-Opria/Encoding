package main

import (
	"crypto/rc4"
	"encoding/base64"
	"fmt"
)

var testExample = []byte{208, 161, 208, 191, 208, 184, 32, 208, 177, 208, 187, 209, 143, 209, 130, 209, 140, 32, 208, 189, 208, 190, 209, 135, 208, 176, 208, 188, 208, 184, 33}

func main() {
	key := "10"
	fmt.Println(testExample)
	encodedString, _ := encode(testExample, key)
	fmt.Println(encodedString)
	decodedSlice, _ := decode(encodedString, key)
	fmt.Println(decodedSlice)
}

func encode(arr []byte, key string) (string, error) {
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Can`t make new cipher", err)
		return "", err
	}
	toRC4 := make([]byte, len(arr))
	cipher.XORKeyStream(toRC4, arr)

	return base64.StdEncoding.EncodeToString([]byte(toRC4)), nil
}

func decode(base64Str string, key string) ([]byte, error) {
	toRC4, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Println("decode base64 error:", err)
		return []byte{}, err
	}
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Can`t make new cipher", err)
		return []byte{}, err
	}
	decodedSlice := make([]byte, len(toRC4))
	cipher.XORKeyStream(decodedSlice, toRC4)

	return decodedSlice, nil
}
