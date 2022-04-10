package main

import "fmt"

type Cat struct {
	name string
}

// (a Cat) -> Cat 전용 메서드
func (a Cat) cry() {
	fmt.Println("야옹~")
}

type Animal struct {
	name string
}

func (a Animal) introduce() {
	fmt.Println("애완동물의 이름은 " + a.name + "입니다.")
}

type Hamster struct {
	// Animal 구조체의 자료형만 선언하면 상속과 같은 역할
	Animal
}

func (h Hamster) introduce() {
	fmt.Println("내 햄스터 이름은 " + h.name + "입니다.")
}

type Dog struct {
	// Animal 구조체의 자료형만 선언하면 상속과 같은 역할
	Animal
}

func (d Dog) introduce() {
	fmt.Println("내 강아지의 이름은 " + d.name + "입니다.")
}

func main() {
	kitty := Cat{"나비"}
	kitty.cry() // 야옹~
	// cry()       // error

	a := Animal{"통통이"}
	mouse := Hamster{}
	puppy := Dog{}

	mouse.Animal.name = "나비"
	puppy.Animal.name = "뭉치"

	a.introduce()     // 내 애완동물의 이름은 꿀꿀이입니다.
	mouse.introduce() // 내 애완동물의 이름은 나비입니다.
	puppy.introduce() // 내 애완동물의 이름은 뭉치입니다.

	mouse.Animal.introduce() // 내 애완동물의 이름은 나비입니다.
	puppy.Animal.introduce() // 내 애완동물의 이름은 뭉치입니다.
}
