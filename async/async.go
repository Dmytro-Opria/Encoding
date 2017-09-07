package main

import (
	"fmt"
	"time"
)

var workChan = make(chan int, 10)

func main() {
	for i := 1; i <= 10; i++ {
		go worker(i)
	}
	printNumber()
}

func worker(inx int) {
	for {
		toPrint(inx)
	}
}

func toPrint(inx int) {
	workChan <- inx
}

func printNumber() {
	for number := range workChan {
		fmt.Println(number)
		time.Sleep(1 * time.Second)
	}
}
