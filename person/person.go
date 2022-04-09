package person

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("SetDetails nico: ", p)
}

func (p *Person) SetDetails2(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println("SetDetails2 nico: ", p)
}
