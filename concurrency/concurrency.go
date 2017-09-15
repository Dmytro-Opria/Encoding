package main

import (
	"fmt"
	"time"
	"sync"
)

var GlobalMap = make(map[string]string, 1)
var ABC = []string{"A","B","C"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	makeGorutine(3)
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

