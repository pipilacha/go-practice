package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	closed  bool
	lock    sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock() // by defering unlock it will execute inmediately after Balance() returns a value

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.closed {
		return 0, false
	}

	new := a.balance + amount
	if new < 0 {
		return 0, false
	}
	a.balance = new

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	old := a.balance
	a.balance = 0

	return old, true
}
