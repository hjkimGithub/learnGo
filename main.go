package main

import "fmt"

func main() {
	a := 2
	fmt.Println(a, &a)
	b := &a
	fmt.Println(a, &a, b, *b)
	a = 12
	fmt.Println(a, &a, b, *b)
}
