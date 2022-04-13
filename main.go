package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

type Student2 struct {
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age: %d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func ConvertType(stringer Stringer) {
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	s := &Student{15}
	PrintAge(s)

	// var stringer Stringer
	// compile time error, impossible type asssertion
	// stringer.(*Student2)

	actor := &Actor{}
	ConvertType(actor)
}
