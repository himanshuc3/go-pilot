package wallet

import (
	"errors"
	"fmt"
)

// NOTE:
// 1. Creating a new type to avoid confusion with int.
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// NOTE:
// 1. `var` used for global variables.
var ErrInsufficientFunds = errors.New("chapri aadmi, no paisa")

// NOTE:
// 1. func (w Wallet) passes the receiver this as a copy of
// the value of the wallet struct.
// 2. Golang supports parameteric polymorphism, which is a form of
// compile time polymorphism uing generics.
func (w *Wallet) Deposit(amount Bitcoin) {
	// NOTE:
	// 1. `&` is used to get the address of a variable and %p is
	// used to print the address in hexadecimal format.
	fmt.Printf("address of balance in deposit %p\n", &w.balance)
	// NOTE:
	// 1. In plain sight, the code accessing pointer variables directly
	// is dereferencing the pointer (which only creates more confusion).
	// 2. Also, how is bitcoin type supporting arithmetic operations?
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// NOTE:
// 1. Error while a new native type in go, is an interface, not a struct.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	// NOTE:
	// 1. %d doing automatic type conversion to int.
	return fmt.Sprintf("%d BTC", b)
}
