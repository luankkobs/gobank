package main

import (
	"fmt"

	"github.com/luankkobs/bank/accounts"
)

func PayBillet(account verifyAccount, valueOfBillet float64) {
	account.Withdrawal(valueOfBillet)
}

type verifyAccount interface {
	Withdrawal(value float64) string
}

func main() {
	denisAccount := accounts.SavingsAccount{}

	denisAccount.Deposit(300)
	PayBillet(&denisAccount, 60)

	fmt.Println(denisAccount.GetBalance())

	luisaAccount := accounts.CheckingAccount{}
	luisaAccount.Deposit(500)

	PayBillet(&luisaAccount, 1000)

	fmt.Println(luisaAccount.GetBalance())

}
