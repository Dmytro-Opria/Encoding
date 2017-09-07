package main

import (
	"testing"
	_"os"
	_"bufio"
	"os/exec"
	"fmt"
	"os"
	"bufio"
)

func TestCreateWCfile(t *testing.T) {
	fileName := "TestFunc.txt"
	value := []byte{}
	CreateWCfile(fileName)

	defer os.Remove(fileName)
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		t.Error("Can not open file", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value = append(value, scanner.Bytes()...)
	}

	cmd := exec.Command("sh","-c","wc -l exec.go > TestCommand.txt")

	err = cmd.Run()

	fmt.Println(err)


	testValue := []byte{}

	fileTest, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		t.Error("Can not open file", err)
	}
	scannerTest := bufio.NewScanner(fileTest)

	for scannerTest.Scan() {
		testValue = append(testValue, scannerTest.Bytes()...)
	}

	if string(value) != string(testValue){
		t.Error("The values are not equal\n", string(testValue),"\n", string(value))
	}
}
