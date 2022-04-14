package main

import (
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

func main() {
	contains := funk.Contains([]int{1, 2, 3, 4, 5}, 1)
	fmt.Println(contains)

	r := funk.Map([]int{1, 2, 3, 4}, func(x int) int {
		return x * 2
	})
	fmt.Println(r)

	r2 := funk.Map([]string{"a", "b", "3"}, strings.ToUpper)
	fmt.Println(r2)
}
