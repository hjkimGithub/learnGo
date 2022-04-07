package main

import (
	"fmt"
)

const (
	// C1 uint = iota + 1
	C1 uint = 1 << iota
	C2
	C3
)

func main() {
	fmt.Println(C1, C2, C3)
}
