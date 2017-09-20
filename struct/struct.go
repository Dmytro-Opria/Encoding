package main

import (
	"fmt"
	"reflect"
	"strings"
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
	fmt.Println(getTagsFromStr(string(field.Tag)))
}

func getTagsFromStr(tagStr string) (res string) {

	splitedTags := strings.Split(tagStr, " ")

	for i, v := range splitedTags {

		tag := strings.Split(v, ":")

		if i+1 == len(splitedTags) {
			res += fmt.Sprintf("Tag: %s\nValue: %s", tag[0], strings.Trim(tag[1], `""`))
			return
		}
		res += fmt.Sprintf("Tag: %s\nValue: %s\n", tag[0], strings.Trim(tag[1], `""`))
	}
	return
}
