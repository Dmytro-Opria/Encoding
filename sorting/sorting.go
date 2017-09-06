package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

var strArr = []string{"Cheese", "Pepperoni", "Black Olives", "Ankit", "Bohra", "Xyz", "AAA", "BBB", "CCC", "DDD", "EEE", "FFF"}
var intArr = []int{1, 45, 62, 32, 31231, 321, 1, 657, 3432, 677554, 333344332234, 6451231}

func main() {
	//strSorting(strArr)
	//fmt.Println(strArr)
	//strReverseSorting(strArr)
	//fmt.Println(strArr)
	//intSorting(intArr)
	//fmt.Println(intArr)
	//intReverseSorting(intArr)
	//fmt.Println(intArr)
	//fmt.Println(intRandomizer(0,100))
	for i := int64(0); i < 1000; i++ {

		fmt.Println(floatRandomizer(i))
	}
}

func strSorting(strArr []string) {
	sort.Slice(strArr, func(i, j int) bool { return strArr[i] < strArr[j] })
}

func strReverseSorting(strArr []string) {
	sort.Slice(strArr, func(i, j int) bool { return strArr[i] > strArr[j] })
}

func intSorting(intArr []int) {
	sort.Sort(sort.IntSlice(intArr))
}

func intReverseSorting(intArr []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(intArr)))
}

func intRandomizer(min, max int) int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func floatRandomizer(min int64) float64{
	source := rand.NewSource(min)
	random := rand.New(source)

	return random.Float64()
}
