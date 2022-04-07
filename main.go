package main

import "fmt"

func main() {
	var a int
	var b int

	// n, err := fmt.Scan(&a, &b)
	// n, err := fmt.Scanf(&a, &b)
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, err)
	}
}
