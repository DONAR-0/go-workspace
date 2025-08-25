package encap

import "fmt"

type BankAccount struct {
	accountHolder string
	balance       float64
}

func NewBankAccount(holder string, balance float64) *BankAccount {
	return &BankAccount{
		accountHolder: holder,
		balance:       balance,
	}
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
		fmt.Println("Deposited:", amount)
	} else {
		fmt.Println("Invalid deposit amount")
	}
}
