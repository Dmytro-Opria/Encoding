package main

import (
	"sync"
	"fmt"
	"time"
)

var (
	wg sync.WaitGroup
	GlobalMap  = sync.Map{}
	ABC = []string{"A", "B", "C"}
)

func main(){
	smap := sync.Map{}
	smap.Store("ten", 10)
	smap.Store(10, "ten")

	val, ok := smap.Load("ten")
	if ok {
		fmt.Println(val)
	}
	//smap.Delete(10)
	val, ok = smap.Load(10)
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Can`t read data")
	}

	val, ok = smap.LoadOrStore("ten", 20)

	if ok {
		fmt.Println(val)
	}
	//makeGorutine(3)
	writeTimeInMap()

}

func makeGorutine(count int) {
	wg.Add(count)
	for i := 0; i < count; i++ {
		timer := time.Now().Add(10 * time.Second)
		go func(n int) {
			for {
				GlobalMap.Store("ABC", ABC[n])
				retVal, ok := GlobalMap.Load("ABC")
				if ok {
					fmt.Println(retVal)
				}

				if time.Now().Unix() > timer.Unix() {
					wg.Done()
					break
				}
			}
		}(i)
	}
	wg.Wait()
}

func writeTimeInMap() {
	timer := time.Now().Add(2 * time.Second)
	//var copy map[interface{}]interface{}
	counter := 0
	for {
		GlobalMap.Store(counter, time.Now().String())

		if time.Now().Unix() > timer.Unix() {
			GlobalMap.Range(func(k, v interface{}) bool {
				//copy[k] = v

				fmt.Println("Key =",k)
				fmt.Println("Value =", v)
				return true
			})
			break
		}
		counter++
	}
}

