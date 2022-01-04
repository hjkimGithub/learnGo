package main

import (
	"fmt"

	"github.com/hjkimGithub/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
