package main

import "fmt"

func main() {
	a := 123
	b := 456
	c := 123456789

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d, %-5d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-5d\n", c, c)

	d := 324.13455
	e := 3.14

	fmt.Printf("%08.2f\n", d)
	fmt.Printf("%08.2g\n", d)
	fmt.Printf("%8.5g\n", d)
	fmt.Printf("%f\n", e)

	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n\\"
	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
