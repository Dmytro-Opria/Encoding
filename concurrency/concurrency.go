package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	GlobalMap  = make(map[string]string, 1)
	GlobalMap2 = map[string]string{"A": "0", "B": "0", "C": "0"}

	ABC = []string{"A", "B", "C"}
	wg  sync.WaitGroup
	mut sync.Mutex

	atomicMap atomic.Value
)

func init() {
	atomicMap.Store(GlobalMap2)
}

func main() {
	//Multiple writing
	//makeGorutine(3)
	//===============
	//Multiple writing/reading with mutex
	//go writeTimeInMap()
	//readMapByGorutine(3)
	//===============
	//Multiple writing/reading with atomic
	go writeTimeInMapAtomic()
	readMapByGorutineAtomic(3)
}

func makeGorutine(count int) {
	wg.Add(count)

	for i := 0; i < count; i++ {
		timer := time.Now().Add(10 * time.Second)
		go func(n int) {
			for {
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

//======Mutex
func readMapByGorutine(count int) {
	wg.Add(count)
	for i := 0; i < count; i++ {
		timer := time.Now().Add(10 * time.Second)
		go func(n int) {
			for {
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

func writeTimeInMap() {
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

//========Atomic

func readMapByGorutineAtomic(count int) {
	wg.Add(count)
	curMap := atomicMap.Load().(map[string]string)

	for i := 0; i < count; i++ {
		timer := time.Now().Add(10 * time.Second)

		go func(n int) {
			ticker := time.NewTicker(100 * time.Millisecond)

		Loop:
			for {
				select {
				case <-ticker.C:
					curMap = atomicMap.Load().(map[string]string)
				default:
					fmt.Println(curMap[ABC[n]])

					if time.Now().Unix() > timer.Unix() {
						break Loop
					}
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func writeTimeInMapAtomic() {
	timer := time.Now().Add(10 * time.Second)
	wg.Add(1)
	for {
		globalMap := map[string]string{"A": "0", "B": "0", "C": "0"}

		for key := range globalMap {
			globalMap[key] = time.Now().String()
		}
		atomicMap.Store(globalMap)
		if time.Now().Unix() > timer.Unix() {
			wg.Done()
			break
		}
	}
	wg.Wait()
}
