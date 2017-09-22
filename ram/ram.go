package main

import (
	"runtime"
	"time"
	"fmt"
)

func main(){
	monitorRuntime()
}

func monitorRuntime() {
	for {
		m := &runtime.MemStats{}
		runtime.ReadMemStats(m)
		fmt.Printf("Idle memory %d from %d\n", m.HeapIdle, m.Sys)
		time.Sleep(5 * time.Second)
	}
}
