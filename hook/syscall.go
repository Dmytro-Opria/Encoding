package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Определяем, куда мы будем выводить сообщения
var writeTo io.Writer = ioutil.Discard
var mutex sync.Mutex

func toggleOutput(c chan os.Signal) {
	for {
		// На этой строчке мы заблокируемся пока не получим сигнал
		<-c

		// Как только мы получили сигнал, необходимо сменить наш writer.
		// Если раньше мы писали в stdout, то меняем на ioutil.Discard
		// и, соответствено, наоборот.
		mutex.Lock()
		if writeTo == os.Stdout {
			writeTo = ioutil.Discard
		} else {
			writeTo = os.Stdout
		}
		mutex.Unlock()
	}
}

func main() {

	fmt.Printf("Process PID : %v\n", os.Getpid())

	c := make(chan os.Signal, 1)
	// В канал `c` попадет сообщение как только получим сигнал SIGUSR1.
	signal.Notify(c, syscall.SIGUSR2)
	go toggleOutput(c)

	// Бесконечный цикл, в котором мы просто увеличиваем счетчик.
	counter := 1
	for {
		mutex.Lock()
		fmt.Fprintln(writeTo, counter)
		mutex.Unlock()
		counter++
		time.Sleep(time.Second)
	}
}