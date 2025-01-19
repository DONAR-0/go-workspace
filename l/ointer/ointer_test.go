package ointer

import (
	"errors"
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		tablewriter.AssertStructGotWant(t, got, want)
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != nil {
			if got.Error() != want.Error() {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, errors.New(""))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw Insufficient Amount", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
