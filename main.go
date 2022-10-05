package main

import (
	"fmt"

	"github.com/luankkobs/bank/accounts"
)

func main() {
	contaExemplo := accounts.CheckingAccount{}
	contaExemplo.Deposit(100)
	fmt.Println(contaExemplo.GetBalance())
}
