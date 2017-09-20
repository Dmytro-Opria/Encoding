package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Field int32 `myTag1:"abc" myTag3:"qwe"`
}

func main() {
	typeA := &A{1}
	field, ok := reflect.TypeOf(typeA).Elem().FieldByName("Field")
	if !ok {
		fmt.Println("Can`t get field")
		return
	}
	fmt.Println(field.Tag.Get("myTag1"))
	fmt.Println(field.Tag.Get("myTag3"))
}
