package person

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SetDetails(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Println("SetDetails nico: ", p)
}

func (p *Person) SetDetails2(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Println("SetDetails2 nico: ", p)
}

func (this Person) HisName() {
	fmt.Printf("his name is %s.\n", this.Name)
}

func (this Person) HisAge() {
	fmt.Printf("his age is %d.\n", this.Age)
}

type getName interface {
	HisName()
}

type getAge interface {
	HisAge()
}

// interface가 내부에 interface를 선언하면 그 선언된 interface의 메소드를 가져오게된다.
type InfoPerson interface {
	getName
	getAge
}
