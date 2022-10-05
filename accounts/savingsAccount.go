package accounts

import (
	"github.com/luankkobs/bank/customers"
)

type SavingsAccount struct {
	AccountHolder                          customers.Holder
	AgencyNumber, AccountNumber, Operation int
	accountBalance                         float64
}

func (c *SavingsAccount) Withdrawal(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= c.accountBalance
	if canWithdraw {
		c.accountBalance -= withdrawalValue
		return "Withdrawal successful"
	} else {
		return "Insufficient funds"
	}
}

func (c *SavingsAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.accountBalance += depositValue
		return "Deposit successful", c.accountBalance
	} else {
		return "Deposit failed.", c.accountBalance
	}
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.accountBalance
}
