package accounts

type IBankAccount interface {
	Withdraw(value float64) (bool, float64)
	Deposit(value float64) (bool, float64)
}
