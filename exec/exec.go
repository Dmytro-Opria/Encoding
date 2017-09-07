package main

import (
	"os/exec"
	"fmt"
	"os"
	"bufio"
	"log"
)


func main(){
	CreateWCfile("execCommand.txt")
	execCommand()
}

func CreateWCfile(fileName string){
	cmd := exec.Command("wc", "-l", "exec.go")

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(cmdReader)
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("File creating eror", err)
		return
	}
	defer file.Close()

	cmd.Start()
	scanner.Scan()
	file.Write(scanner.Bytes())
}

func execCommand() {
	cmd := exec.Command("wc", "-l", "exec.go", ">", "test.txt")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

