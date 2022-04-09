package main

import "fmt"

func main() {
	foods := []string{"potato", "pizza", "pasta"}
	// for i := 0; i < len(foods); i++ {
	// 	fmt.Println(foods[i])
	// }
	fmt.Printf("%v\n", foods)
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)
}
