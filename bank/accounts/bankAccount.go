package accounts

import (
	"bank/owner"
	"fmt"
)

type BankAccount struct {
	AccountNumber int
	AgencyNumber  int
	Owner         *owner.Owner
	Balance       float64
}

func (this *BankAccount) Withdraw(value float64) (bool, float64) {
	if value < 0 {
		fmt.Println("Valor invalido")
		return false, this.Balance
	}

	if value >= this.Balance {
		fmt.Println("Saldo insuficiente")
		return false, this.Balance
	}

	this.Balance -= value
	return true, this.Balance
}

func (this *BankAccount) Transfer(value float64, account *BankAccount) bool {
	if value < 0 || value > this.Balance {
		return false
	}

	this.Withdraw(value)
	account.Deposit(value)
	return true
}

func (this *BankAccount) Deposit(value float64) (bool, float64) {
	if value <= 0 {
		return false, this.Balance
	}

	this.Balance += value
	return true, this.Balance
}
