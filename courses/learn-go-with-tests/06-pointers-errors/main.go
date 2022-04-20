package main

import (
	"errors"
	"fmt"
)

type SomeDigitalCoin int

// This implements fmt.Stringer built-in interface
func (c SomeDigitalCoin) String() string {
	return fmt.Sprintf("%d SDC", c)
}

type Wallet struct {
	balance SomeDigitalCoin
}

func (w *Wallet) Deposit(amount SomeDigitalCoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount SomeDigitalCoin) error {
	if w.balance < amount {
		return errors.New("cannot withdraw from wallet")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() SomeDigitalCoin {
	return w.balance
	// Equivalent to
	// return (*w).balance
}
