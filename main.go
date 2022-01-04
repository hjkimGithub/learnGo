package main

import (
	"fmt"

	"github.com/hjkimGithub/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	// err := account.Withdraw(15)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(account)
}
