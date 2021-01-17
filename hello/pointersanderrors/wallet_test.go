package pointersanderrors

import "testing"

func TestDeposit(t *testing.T) {
	wallet := Wallet{}

	depositAmount := Bitcoin(10)
	wallet.Deposit(depositAmount)

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw when there is enough money", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}

		if err != nil {
			t.Errorf("There was an error but there shouldn't have been one")
		}
	})

	t.Run("Withdraw when there isn't enough money", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(20))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}

		if err == nil {
			t.Errorf("Trying to withdraw when funds aren't available should have thrown an error")
		}
	})
}
