package main

import (
	"fmt"
	"reflect"
	"strings"
)

type A struct {
	Field int32 `myTag1:"abc" myTag3:"qwe"`
	Field1 int32 `myTag2:"zxc" myTag4:"dfg"`
	Field3 int32 `myTag5:"jkl" myTag7:"bnm"`
}

func main() {
	typeA := A{}

	reflectValue := reflect.ValueOf(typeA)
	typ := reflectValue.Type()

	for i := 0; i < typ.NumField(); i++ {

		field := typ.Field(i)

		if field.PkgPath != "" {
			continue
		}

		fmt.Println(getTagsFromStr(string(field.Tag)))
		fmt.Println("========================================")
	}
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
