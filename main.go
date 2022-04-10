package main

import "fmt"

func main() {
	// 	str1 := "Hello\n'World'\t\n"
	// 	str2 := `Hello
	// 'World'
	// `
	// 	fmt.Println(str1)
	// 	fmt.Println()
	// 	fmt.Println(str2)

	str3 := "Hello 월드!"
	runes := []rune(str3)

	fmt.Println(len(str3), len(runes))

	for i := 0; i < len(runes); i++ {
		// fmt.Printf("%T, %d, %c\n", str3[i], str3[i], str3[i])
		fmt.Printf("%T, %d, %c\n", runes[i], runes[i], runes[i])
	}

	fmt.Println()

	for _, v := range str3 {
		fmt.Printf("%T, %d, %c\n", v, v, v)
	}
}
