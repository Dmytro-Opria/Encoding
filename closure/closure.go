package main

import "fmt"

func GetFib() func() int64{
	previousInt, startInt := int64(0), int64(1)

	return func() int64{
		if startInt == 1 {
			previousInt = startInt
			startInt += 1
			return 1
		}

		previousInt, startInt =  startInt, startInt + previousInt

		return startInt
	}

}

func main(){

	fib := GetFib()

	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())

	newFib := GetFib()

	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())
	fmt.Println(newFib())

}
