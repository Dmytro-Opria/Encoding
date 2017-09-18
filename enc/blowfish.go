package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

var testExample2 = []byte{208, 161, 208, 191, 208, 184, 32, 208, 177, 208, 187, 209, 143, 209, 130, 209, 140, 32, 208, 189, 208, 190, 209, 135, 208, 176, 208, 188, 208, 184, 33}

func main() {
	key := "100"
	fmt.Println(testExample2)
	encodedString, length, _ := encodeBlowFish(testExample2, key)
	fmt.Println(encodedString)
	decodedSlice, _ := decodeBlowFish(encodedString, length, key)
	fmt.Println(decodedSlice)
}

func encodeBlowFish(arr []byte, key string) (string, int, error) {
	cipher, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Can`t make new cipher", err)

		return "", 0, err
	}
	toBlowFish := []byte{}

	for _, v := range makeSlice(arr) {
		part := make([]byte, 8)

		cipher.Encrypt(part, v)

		toBlowFish = append(toBlowFish, part...)
	}
	return base64.StdEncoding.EncodeToString([]byte(toBlowFish)), len(arr), nil
}

func decodeBlowFish(base64Str string, length int, key string) ([]byte, error) {
	toBlowFish, err := base64.StdEncoding.DecodeString(base64Str)

	if err != nil {
		fmt.Println("decode base64 error:", err)

		return []byte{}, err
	}
	cipher, err := blowfish.NewCipher([]byte(key))

	if err != nil {
		fmt.Println("Can`t make new cipher", err)

		return []byte{}, err
	}

	decodedSlice := []byte{}

	for _, v := range makeSlice(toBlowFish) {
		part := make([]byte, 8)

		cipher.Decrypt(part, v)

		decodedSlice = append(decodedSlice, part...)
	}
	return decodedSlice[:length], nil
}

func makeSlice(slice []byte) (val [][]byte) {
	for i := 0; i < len(slice); i += 8 {

		if i+8 > len(slice) {

			sl := make([]byte, 8)

			copy(sl, slice[i:])

			val = append(val, sl)
			break
		}
		val = append(val, slice[i:i+8])
	}
	return
}
