package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}

type User2 struct {
	A int8
	B int
	C int8
	D int
	E int8
}

type User3 struct {
	A int8
	B int8
	C int8
	D int
	E int
}

func main() {
	user2 := User2{1, 2, 3, 4, 5}
	user3 := User3{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user2), unsafe.Sizeof(user3))
}
