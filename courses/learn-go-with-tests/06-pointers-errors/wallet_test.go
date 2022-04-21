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

	assertError := func(t testing.TB, result, expected error) {
		t.Helper()
		if result == nil {
			// New: stop test here if this gets called!
			t.Fatal("should be an error")
		}

		if result != expected {
			t.Errorf("Result: %q Expected: %q", result.Error(), expected)
		}
	}

	assertNoError := func(t testing.TB, result error) {
		t.Helper()
		if result != nil {
			t.Fatal("should not be an error")
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
		err := w.Withdraw(SomeDigitalCoin(10))
		expected := SomeDigitalCoin(10)
		assertNoError(t, err)
		assertBalance(t, w, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := SomeDigitalCoin(20)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(SomeDigitalCoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, startingBalance)
	})
}
