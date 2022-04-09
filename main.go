package main

import (
	"fmt"

	"github.com/hjkimGithub/learnGo/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println(nico)
	nico.SetDetails2("nico", 12)
	fmt.Println(nico)
}
