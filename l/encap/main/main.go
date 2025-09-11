package main

import (
	"fmt"

	en "github.com/donar-0/go-workspace/l/encap"
)

func main() {
	account := en.NewBankAccount("Alice", 1000)
	fmt.Println("Current Balance:", account.GetBalance())
	account.Deposit(500)
	fmt.Println("Update Balance:", account.GetBalance())

	emp := en.Employee{}
	emp.SetName("John Doe")
	emp.SetAge(25)
	fmt.Println("Employee Name:", emp.GetName())
	fmt.Println("Employee Age:", emp.GetAge())

	my_account := en.NewAccount(1000)
	my_account.Withdraw(2000)
	fmt.Println("Curent balance", my_account.GetBalance())
}
