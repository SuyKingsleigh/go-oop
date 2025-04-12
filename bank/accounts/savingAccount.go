package accounts

import (
	"bank/owner"
	"fmt"
)

type SavingAccount struct {
	AccountNumber int
	AgencyNumber  int
	Owner         *owner.Owner
	Balance       float64
}

func (this *SavingAccount) Withdraw(value float64) (bool, float64) {
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

func (this *SavingAccount) Transfer(value float64, account *SavingAccount) bool {
	if value < 0 || value > this.Balance {
		return false
	}

	this.Withdraw(value)
	account.Deposit(value)
	return true
}

func (this *SavingAccount) Deposit(value float64) (bool, float64) {
	if value <= 0 {
		return false, this.Balance
	}

	this.Balance += value
	return true, this.Balance
}
