package main

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"os/exec"
	"log"
)

func TestCreateWCfile(t *testing.T) {
	fileName := "TestFunc.txt"
	testValue := []byte{}
	CreateWCfile(fileName)

	defer os.Remove(fileName)
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("OpenFile Error", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		testValue = append(testValue, scanner.Bytes()...)
	}

	fileNameCom := "TestCommand.txt"

	cmd := exec.Command("wc", "-l", "exec.go", ">", fileNameCom)
	log.Printf("Running command and waiting for it to finish...")
	cmd.Start()
}
