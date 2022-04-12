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

func (this Person) HisName() {
	fmt.Printf("his name is %s.\n", this.name)
}

func (this Person) HisAge() {
	fmt.Printf("his age is %d.\n", this.age)
}

type GetName interface {
	HisName()
}

type GetAge interface {
	HisAge()
}

// interface가 내부에 interface를 선언하면 그 선언된 interface의 메소드를 가져오게된다.
type InfoPerson interface {
	GetName
	GetAge
}
