package pointers

import "errors"

var ErrInsufficientFounds = errors.New("insufficient funds")

// Bitcoin currency type
type Bitcoin int

// Wallet holds a balance of bitcoins
type Wallet struct {
	balance Bitcoin
}

// Balance of a wallet
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit n money into the wallet
func (w *Wallet) Deposit(n Bitcoin) {
	w.balance += n
}

// Withdraw n money of the wallet
func (w *Wallet) Withdraw(n Bitcoin) error {
	if n > w.balance {
		return ErrInsufficientFounds
	}
	w.balance -= n
	return nil
}
