package main

import "fmt"

type opFunc func(a, b int) int

func getOp(op string) func(int, int) int {
	if op == "+" {
		return func(i1, i2 int) int { return i1 + i2 }
	} else if op == "*" {
		return func(i1, i2 int) int { return i1 * i2 }
	} else {
		return nil
	}
}

func main() {
	var op func(int, int) int
	// cannot use initial :=
	// op := func(int, int) int
	op = getOp("*")

	var result = op(3, 4)
	fmt.Println(result)
}
