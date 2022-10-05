package main

import (
	"fmt"

	"github.com/luankkobs/bank/accounts"
)

func main() {
	denisAccount := accounts.SavingsAccount{}

	denisAccount.Deposit(300)
	denisAccount.Withdrawal(200)

	fmt.Println(denisAccount.GetBalance())

}
