package main

import (
	"fmt"

	"github.com/hjkimGithub/learnGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "First")
	fmt.Println(word)
	dictionary.Update(word, "Helloo")
	word, _ = dictionary.Search(word)
	fmt.Println(word)
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	fmt.Println(err)
	// err := dictionary.Update("bye", "BYEBYE")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
