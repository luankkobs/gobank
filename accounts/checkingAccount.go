package accounts

import (
	"github.com/luankkobs/bank/customers"
)

type CheckingAccount struct {
	AccountHolder  customers.Holder
	AgencyNumber   int
	AccountNumber  int
	accountBalance float64
}

func (c *CheckingAccount) Withdrawal(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= c.accountBalance
	if canWithdraw {
		c.accountBalance -= withdrawalValue
		return "Withdrawal successful"
	} else {
		return "Insufficient funds"
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.accountBalance += depositValue
		return "Deposit successful", c.accountBalance
	} else {
		return "Deposit failed.", c.accountBalance
	}
}

func (c *CheckingAccount) Transfer(transferValue float64, destinationAccount *CheckingAccount) bool {
	if transferValue <= c.accountBalance && transferValue > 0 {
		c.accountBalance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.accountBalance
}
