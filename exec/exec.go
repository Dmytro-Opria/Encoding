package main

import (
	"os/exec"
	"os"
	"bufio"
)


func main(){
	CreateWCfile("execCommand.txt")
}

func CreateWCfile(fileName string)(err error){
	cmd := exec.Command("wc", "-l", "exec.go")

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		os.Exit(1)
		return
	}
	scanner := bufio.NewScanner(cmdReader)
	file, err := os.Create(fileName)

	if err != nil {
		return
	}
	defer file.Close()

	cmd.Start()

	value := []byte{}

	for scanner.Scan() {
		value = append(value, scanner.Bytes()...)
	}

	file.Write(value)

	return nil
}
