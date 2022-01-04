package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("CANNOT WITHDRAW")

// NewAccount create new accouts
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Change Owner of the Account
func (a *Account) ChangeOwner(newOwners string) {
	a.owner = newOwners
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// internal method
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s Account. \n Has: ", a.Balance())
}
