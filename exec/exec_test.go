package main

import (
	"testing"
	"os"
	"bufio"
	"os/exec"
)

func TestCreateWCfile(t *testing.T) {
	fileName := "TestFunc.txt"
	testValue := []byte{}
	CreateWCfile(fileName)

	defer os.Remove(fileName)
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		t.Error("Can not open file", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		testValue = append(testValue, scanner.Bytes()...)
	}

	fileNameCom := "TestCommand.txt"

	fileCom, _ := os.Create(fileNameCom)

	cmd := exec.Command("wc", "-l", "exec.go", ">", fileNameCom)

	cmd.Stdout = fileCom

	/*if  strings.Contains(string(testValue), string(cmd)){
		t.Error("The values are not equal\n", string(testValue),"\n", string(cmd))
	}*/
}
