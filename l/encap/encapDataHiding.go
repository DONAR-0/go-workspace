package encap

import (
	"fmt"
)

// Account struct with previous balance
type Account struct {
	balance float64
}

func NewAccount(initBalance float64) *Account {
	return &Account{balance: initBalance}
}

// Private Method for withdrawal validation
func (a *Account) validateWithWithDrawal(amount float64) bool {
	return amount > 0 && amount <= a.balance
}

// Public Method
func (a *Account) Withdraw(amount float64) {
	if a.validateWithWithDrawal(amount) {
		a.balance -= amount
		fmt.Println("Withdrawal Sucessful:", amount)
	} else {
		fmt.Println("Insufficient balance or invalid amount")
	}
}

// Getter balance
func (a *Account) GetBalance() float64 {
	return a.balance
}
