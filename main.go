package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) add(a int) {
	p.X += a
	p.Y += a
}

func (p Point) mul(a int) {
	p.X *= a
	p.Y *= a
}

type Square struct {
	Width  float32
	Height float32
}

type Triangle struct {
	Width  float32
	Height float32
}

func (s Square) Area() float32 {
	return s.Width * s.Height
}

func (t Triangle) Area() float32 {
	return t.Width * t.Height / 2
}

type Line struct {
	Length float32
}

type Figure interface {
	Area() float32
}

type Person struct {
	Name string
	Age  int
}

func (this Person) hisName() {
	fmt.Printf("his name is %s.\n", this.Name)
}

func (this Person) hisAge() {
	fmt.Printf("his age is %d.\n", this.Age)
}

type GetName interface {
	hisName()
}

type GetAge interface {
	hisAge()
}

// interface가 내부에 interface를 선언하면 그 선언된 interface의 메소드를 가져오게된다.
type InfoPerson interface {
	GetName
	GetAge
}

func main() {
	p := Point{3, 4}
	p.add(10)  // 적용됨 13, 14
	p.mul(100) // 적용 불가 13, 14 그대로
	fmt.Println(p)

	s := Square{2, 5}
	t := Triangle{2, 5}
	fmt.Println(s.Area())
	fmt.Println(t.Area())

	var s1, t1 Figure
	s1 = Square{2, 5}
	t1 = Triangle{2, 5}
	fmt.Println(s1.Area())
	fmt.Println(t1.Area())

	f := make([]Figure, 0)
	f = append(f, Square{2, 5})
	f = append(f, Triangle{2, 5})
	// struct가 interface형으로 구현될려면 interface에 정의된 메소드가 struct에 존재해야한다.
	// f = append(f, Line{Length: 5})
	for _, value := range f {
		fmt.Println(value.Area())
	}

	var i InfoPerson
	i = Person{"Kukaro", 27}
	i.hisName()
	i.hisAge()
}
