package main

import (
	"fmt"
	"strings"

	collectionF "github.com/hjkimGithub/learnGo/collectionFunctionEx"
)

func add10(n int) int {
	return n + 10
}

func main() {
	strs1 := []int{1, 2, 3}
	fmt.Println(collectionF.Map(strs1, add10))

	strs2 := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(collectionF.Filter(strs2, func(s string) bool {
		return strings.Contains(s, "e")
	}))

}
