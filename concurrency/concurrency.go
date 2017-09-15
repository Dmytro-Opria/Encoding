package main

import (
	"fmt"
	"time"
	"sync"
)

var GlobalMap = make(map[string]string, 1)
var GlobalMap2 = map[string]string{"A":"0","B":"0","C":"0"}

var ABC = []string{"A","B","C"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	//makeGorutine(3)
	go writeTimeInMap()
	readMapByGorutine(3)
}

func makeGorutine(count int) {
	wg.Add(count)

	for i := 0; i < count ; i++ {
		timer := time.Now().Add(10 * time.Second)
		go func(n int) {
			for  {
				mut.Lock()
				GlobalMap["ABC"] = ABC[n]
				fmt.Println(ABC[n])
				mut.Unlock()

				if time.Now().Unix() > timer.Unix() {
					wg.Done()
					break
				}
			}
		}(i)
	}
	wg.Wait()
}

func readMapByGorutine(count int){
	wg.Add(count)
	for i := 0; i < count ; i++ {
		timer := time.Now().Add(10 * time.Second)
		go func(n int) {
			for  {
				mut.Lock()
				fmt.Println(GlobalMap2[ABC[n]])
				mut.Unlock()

				if time.Now().Unix() > timer.Unix() {
					wg.Done()
					break
				}
			}
		}(i)
	}
	wg.Wait()
}

func writeTimeInMap(){
	timer := time.Now().Add(10 * time.Second)
	wg.Add(1)
	for {
		for key := range GlobalMap2 {
			mut.Lock()
			GlobalMap2[key] = time.Now().String()
			mut.Unlock()
		}
		if time.Now().Unix() > timer.Unix() {
			wg.Done()
			break
		}
	}
	wg.Wait()
}

