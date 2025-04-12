package main

import (
	"bank/accounts"
	"bank/owner"
	"fmt"
)

func main() {
	suy := owner.Owner{
		Name:      "Suyan Moriel",
		CPF:       "123.456.789-00",
		BirthDate: "01/0/1970",
	}

	account := accounts.BankAccount{
		AccountNumber: 12345,
		AgencyNumber:  123,
		Owner:         &suy,
		Balance:       1000.00,
	}

	account.Deposit(100.)
	fmt.Printf("Saldo final %.2f da conta de %s\n", account.Balance, account.Owner.Name)

}
