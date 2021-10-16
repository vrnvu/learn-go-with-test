package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assert.Equal(t, want, got)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assert.NotErrorIs(t, err, ErrInsufficientFounds)

		got := wallet.Balance()
		want := Bitcoin(10)
		assert.Equal(t, want, got)
	})
	t.Run("Withdraw insufficient founds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assert.Error(t, err)
	})
}
