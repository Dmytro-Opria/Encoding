package main

import (
	"os"
	"fmt"
)

func main() {
	progArgs := os.Args
	fmt.Println(progArgs)
	fmt.Println(progArgs[1:])
}
