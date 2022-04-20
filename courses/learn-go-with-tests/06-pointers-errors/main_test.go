package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, expected SomeDigitalCoin) {
		t.Helper()
		result := w.Balance()

		if result != expected {
			t.Errorf("Result: %s Expected: %s", result, expected)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("withdraw should return an error")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(SomeDigitalCoin(10))
		expected := SomeDigitalCoin(10)
		assertBalance(t, w, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{balance: SomeDigitalCoin(20)}
		w.Withdraw(SomeDigitalCoin(10))
		expected := SomeDigitalCoin(10)
		assertBalance(t, w, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := SomeDigitalCoin(20)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(SomeDigitalCoin(100))
		assertError(t, err)
		assertBalance(t, w, startingBalance)
	})
}
