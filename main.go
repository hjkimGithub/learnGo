package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3
	fmt.Println(g, h, f)

	var i float32 = 1234.523
	var j float32 = 3456.123
	var k float32 = i * j
	var l float32 = k * 3

	fmt.Println(i, j, k, l)
}
