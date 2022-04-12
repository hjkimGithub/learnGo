package main

import (
	"fmt"

	"github.com/hjkimGithub/learnGo/person"
	"github.com/hjkimGithub/learnGo/shape"
)

func main() {
	p := shape.Point{3, 4}
	p.Add(10)  // 적용됨 13, 14
	p.Mul(100) // 적용 불가 13, 14 그대로
	fmt.Println(p)

	s := shape.Square{2, 5}
	t := shape.Triangle{2, 5}
	fmt.Println(s.Area())
	fmt.Println(t.Area())

	var s1, t1 shape.Figure
	s1 = shape.Square{2, 5}
	t1 = shape.Triangle{2, 5}
	fmt.Println(s1.Area())
	fmt.Println(t1.Area())

	f := make([]shape.Figure, 0)
	f = append(f, shape.Square{2, 5})
	f = append(f, shape.Triangle{2, 5})
	// struct가 interface형으로 구현될려면 interface에 정의된 메소드가 struct에 존재해야한다.
	// f = append(f, Line{Length: 5})
	for _, value := range f {
		fmt.Println(value.Area())
	}

	i := person.Person{}
	i.SetDetails2("KAYLE", 20)
	i.HisName()
	i.HisAge()
}
