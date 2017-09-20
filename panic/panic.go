package main

import "fmt"

var counter = 0

func main() {
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
	fmt.Println(funcWithPanic())
}

func funcWithPanic() (val int8) {
	defer func() {

		counter++

		if r := recover(); r != nil {
			val = 1
			return
		}
		val = 0
	}()

	if counter%2 == 0 {
		panic("Just for test")
	}
	return val
}
