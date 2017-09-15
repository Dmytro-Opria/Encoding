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
		timer := time.Now().Second() + 10
		go func() {
			for  {
				mut.Lock()
				GlobalMap["ABC"] = ABC[i-1]
				fmt.Println(ABC[i-1])
				mut.Unlock()

				if time.Now().Second() > timer {
					wg.Done()
					break
				}
			}
		}()
	}
	wg.Wait()
}

