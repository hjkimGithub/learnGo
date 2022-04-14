package main

import (
	"container/list"
	"fmt"

	"github.com/hjkimGithub/learnGo/datastructure"
)

func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	queue := datastructure.NewQueue()
	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v2 := queue.Pop()
	for v2 != nil {
		fmt.Printf("%d ->\n", v)
		v2 = queue.Pop()
	}
}
