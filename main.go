package main

import "fmt"

type Stringer interface {
	String() string
}

// type Student struct implements Stringer
type Student struct { // duck typing
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d 살, %s 라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	stringer = student

	fmt.Println(stringer.String())
}
