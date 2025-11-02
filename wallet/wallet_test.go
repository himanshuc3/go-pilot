package wallet

import "testing"

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got == nil {
		// NOTE:
		// 1. `t.Fatal` is used to stop the execution of the test if the condition is true,
		// which apparently is not the case with t.Errorf.
		t.Fatal("didn't get an error but expected one")
	}

	if got != expected {
		t.Errorf("got %q, expected %q", got.Error(), expected)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got an error but didn't want one: %s", got)
	}
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

		// NOTE:
		// 1. nil is equivalent to null in other languages.
		// 2. interfaces can be nil as well.
		// if err == nil {
		// 	t.Errorf("wanted an error but didn't get one")
		// }
	})

}
