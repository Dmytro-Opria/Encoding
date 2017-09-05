package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/textproto"
	"os"
	"strconv"
	"strings"
)

type A struct {
	A int16  `json:"a_field,omitempty"`
	B uint64 `json:"b_field,omitempty"`
	C int    `json:"c_field,omitempty"`
	S string `json:"s_field,omitempty"`
}
type BiteA struct {
	A int16
	B uint64
	C int64
}

func main() {
	pathJson := "test_arr1.json"
	writeArr(pathJson)
	//fmt.Println(readLines(pathJson))

	pathBinary := "test_binary.txt"

	writeBinarySlice(pathBinary)

	binaryDecode(pathBinary)
}

func writeArr(path string) {
	testSlice := getSlice()
	for i, v := range testSlice {
		if i == 0 {
			writeFirstBreak(path)
		}
		val, _ := json.Marshal(v)

		if i+1 == len(testSlice) {
			writeLine(val, path, true)
			writeLastBreak(path)
		} else {
			writeLine(val, path, false)
		}
	}
}

func writeFirstBreak(path string) {
	file, _ := os.Create(path)

	defer file.Close()

	file.WriteString("[\n")
}

func writeLastBreak(path string) {
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	defer file.Close()

	file.WriteString("\n]")
}

func readLines(path string) (Arr []A, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	reader := textproto.NewReader(bufio.NewReader(file))

	for {
		line, err := reader.ReadLine()
		if err != nil {
			if line != "[" {
				break
			}
		}
		a := A{}
		json.Unmarshal([]byte(strings.TrimSuffix(line, ",")), &a)
		Arr = append(Arr, a)
	}
	return
}

func writeLine(line []byte, path string, last bool) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if last {
		file.WriteString(string(line))
	} else {
		file.WriteString(string(line) + ",\n")
	}
	return nil
}

/*
func serializeSlice(fileName string, Test []A) (err error) {

	buf := []byte{}
	buf = append(buf, []byte("[\n\t")...)

	for i, a := range Test {
		jsonTest, _ := json.Marshal(a)
		buf = append(buf, jsonTest...)
		if i+1 == len(Test){
			buf = append(buf, []byte("\n")...)
			buf = append(buf, []byte("]")...)
		}else{
			buf = append(buf, []byte(",\n\t")...)
		}
	}
	err = writeByte(fileName, buf)

	return nil
}

func writeByte(fileName string, byte []byte)(err error){
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("os.OpenFile\n %s", err)
		return
	}

	_, err = file.Write(byte)

	if err != nil {
		fmt.Printf("file.Write\n %s", err)
		return
	}
	return nil
}
*/

func getSlice() []A {
	ASlice := make([]A, 10)
	for i := range ASlice {
		ASlice[i].A = int16(i+1) * 100
		ASlice[i].B = uint64(i+1) * 2000
		ASlice[i].C = 1111111111
		ASlice[i].S = "Testing_string_#" + strconv.Itoa(i+1)
	}
	return ASlice
}

func writeBinarySlice(path string) (err error){
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	for _, a := range getSlice() {
		file.Write(binaryEncode(a))
	}

	return
}

func binaryEncode(test A) ([]byte) {
	curStruct := BiteA{int16(test.A), uint64(test.B), int64(test.C)}
	curStr := []byte(test.S)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, curStruct)
	if err != nil {
		panic(err)
	}
	length := make([]byte, 8)

	binary.LittleEndian.PutUint64(length, uint64(len(curStr)))

	buf.Write(length)

	buf.Write(curStr)

	fmt.Println(buf.Bytes())

	return buf.Bytes()
}

func binaryDecode(path string) (ASlice []A){

	f, err := os.Open(path)

	if err != nil {
		return
	}

	for {
		b := make([]byte, 18)

		_, err := f.Read(b)

		val := A{}

		val.A = int16(binary.BigEndian.Uint16(b[0:2]))
		val.B = binary.BigEndian.Uint64(b[2:10])
		val.C = int(binary.BigEndian.Uint64(b[10:18]))

		if err != nil {
			break
		}

		lenStr := make([]byte, 8)

		_, err = f.Read(lenStr)

		valueLength := int(binary.LittleEndian.Uint64(lenStr))

		str := make([]byte, valueLength)

		_, err = f.Read(str)

		val.S = string(str)

		fmt.Println(val)

		ASlice = append(ASlice, val)

		if err != nil {
			break
		}

	}
	return
}

