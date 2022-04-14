package main

import (
	"fmt"

	"github.com/hjkimGithub/learnGo/filerw"
)

const filename string = "data.txt"

func main() {
	line, err := filerw.ReadFile(filename)
	if err != nil {
		err = filerw.WriteFile(filename, "switch a text")
		if err != nil {
			fmt.Println("파일 생성 실패", err)
			return
		}
		line, err = filerw.ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기 실패", err)
			return
		}
	}
	fmt.Println(line)
}
