package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOp(op string) func(int, int) int {
	// return op == "+" ? add : op == "*" ? mul : nil
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
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
