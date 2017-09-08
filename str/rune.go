package main

import "fmt"

func main() {
	str := "абвгдеёжзийклмнопрст"
	runned := []rune(str)
	trimmed := runned[7:]
	fmt.Println(string(trimmed))
}
