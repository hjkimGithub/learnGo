package main

import (
	"fmt"
	"math"
)

const epsilon = 0.00001

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func improvedEqual(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// a, _ := new(big.Float).SetString("0.1")
	// b, _ := new(big.Float).SetString("0.2")
	// c, _ := new(big.Float).SetString("0.3")
	// new(big.Float).Add()
	// c.Cmp(b)

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.000000000004
	b = 0.000000000002
	c = 0.000000000007

	fmt.Printf("%g == %g : %v \n", c, a+b, equal((a+b), c))
	fmt.Printf("%g == %g : %v \n", c, a+b, improvedEqual((a+b), c))
}
