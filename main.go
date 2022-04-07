package main

import (
	"fmt"
)

func getMyAge() int {
	return 22
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "RED"
	case Blue:
		return "BLUE"
	case Green:
		return "GREEN"
	case Yellow:
		return "YELLOW"
	default:
		return "UNDEFINED"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("m,tu")
	case "wednesday", "thursday", "friday":
		fmt.Println("w, th, f")
	}

	temp := 18

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("no outside")
	case temp >= 10 && temp < 20:
		fmt.Println("slightly cold")
	case temp >= 15 && temp < 25:
		fmt.Println("outside")
	default:
		fmt.Println("warm")
	}

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("teen")
	case 33:
		fmt.Println("pair 3")
	default:
		fmt.Printf("age: %d\n", age)
	}
	// age gone
	// fmt.Println(age)

	// switch age2 := getMyAge(); true {
	switch age2 := getMyAge(); {
	case age2 < 10:
		fmt.Println("Child")
	case age2 < 20:
		fmt.Println("Teen")
		// break
	case age2 < 30:
		fmt.Println("20s")
		// fallthrough
	default:
		fmt.Println("My age: ", age2)
	}

	fmt.Println("My favorite color is ", colorToString(getMyFavoriteColor()))
}
