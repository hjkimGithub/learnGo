package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

// better memory usage
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	str1 := "Hello World!"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)

	var str3 string = "Hello world"
	var slice []byte = []byte(str3)

	// string: immutable
	// str3[2] = 'a'
	slice[2] = 'a'

	stringHeader3 := (*reflect.StringHeader)(unsafe.Pointer(&str3))
	sliceHeader3 := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Println(stringHeader3.Data)
	fmt.Println(sliceHeader3.Data)

	var str4 string = "Hello World!!"
	fmt.Println(ToUpper1(str4))
	fmt.Println(ToUpper2(str4))
}
