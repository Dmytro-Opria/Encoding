package main

import (
	"os/exec"
	"os"
)


func main(){
	CreateWCfile("execCommand.txt")
}

func CreateWCfile(fileName string)(err error){
	cmd := exec.Command("wc", "-l", "exec.go")

	file, err := os.Create(fileName)

	cmd.Stdout = file

	cmd.Run()

	if err != nil {
		return
	}
	defer file.Close()

	return nil
}
