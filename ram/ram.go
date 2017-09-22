package main

import (
	_"runtime"
	_"time"
	"fmt"
	"os"
	"bytes"
	"strings"
	"time"
)

func main(){
	monitorMemory()
}
/*
func monitorRuntime() {
	for {
		m := &runtime.MemStats{}
		runtime.ReadMemStats(m)
		fmt.Printf("Idle memory %d from %d\n", m.HeapIdle, m.Sys)
		time.Sleep(5 * time.Second)
	}
}*/

func memoryQuery() (memTotal, memFree string){
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("Can`t execute", err)
		return
	}

	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)


	fields := strings.Split(buf.String(), "\n")
	for _, v := range fields {
		if strings.Contains(v, "MemTotal") {
			slice := strings.Split(v, " ")
			memTotal = slice[len(slice)-2] + " " + slice[len(slice)-1]
		}
		if strings.Contains(v, "MemFree") {
			slice := strings.Split(v, " ")
			memFree = slice[len(slice)-2] + " " + slice[len(slice)-1]
		}
	}
	return
}

func monitorMemory() {
	for {
		memTotal, memFree := memoryQuery()
		//fmt.Printf("Memory Total = %s, Memory Free = %s\n", memTotal, memFree)
		fmt.Printf("\rMemory Total = %s, Memory Free = %s", memTotal, memFree)
		time.Sleep(1 * time.Second)
	}
}
