package main

import (
	_ "io"
	"time"
	"math/rand"
	"fmt"
)

type Reader struct {
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789")

func main(){
	r := Reader{}
	b := make([]byte, 10)
	r.Read(b)
	fmt.Println(string(b))
}

func(r Reader) Read(p []byte) (n int, err error){
	for i := 0; i < 10; i++ {
		p[i] = []byte(string(letterRunes[intRandomizer(0,len(letterRunes), i)]))[0]
	}
	return n, nil
}

func intRandomizer(min, max, seed int) int {
	rand.Seed(time.Now().Unix() + int64(seed))
	return rand.Intn(max-min) + min
}
