package pointersanderrors

import (
	"errors"
	"fmt"
)

// Bitcoin is a digital currency
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// ErrInsufficientFunds is thrown when the wallet has insufficient money inside to perform a given operation
var ErrInsufficientFunds = errors.New("Balance unavailable")

// Wallet represents a wallet that can hold money
type Wallet struct {
	balance Bitcoin
}

// Deposit puts money in the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw removes money from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

// Balance tells you how much money is in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
