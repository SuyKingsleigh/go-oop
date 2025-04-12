package main

import (
	"bank/accounts"
	"bank/owner"
	"fmt"
)

func pay(from accounts.IBankAccount, to accounts.IBankAccount, amount float64) bool {
	from.Withdraw(amount)
	to.Deposit(amount)
	return true
}

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

	blenda := owner.Owner{
		Name:      "Blenda Moriel",
		CPF:       "999.999.999-99",
		BirthDate: "01/0/1970",
	}
	account2 := accounts.SavingAccount{
		AccountNumber: 999,
		AgencyNumber:  888,
		Owner:         &blenda,
		Balance:       1200.0,
	}
	pay(&account2, &account, 100.0)

	fmt.Printf("Saldo final %.2f da conta de %s\n", account.Balance, account.Owner.Name)
	fmt.Printf("Saldo final %.2f da conta de %s\n", account2.Balance, account2.Owner.Name)

}
