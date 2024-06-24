package account

func Open2(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

func (a *Account) Balance2() (int64, bool) {
	a.lock.Lock()

	n, success := int64(0), false
	if !a.closed {
		n, success = a.balance, true
	}
	defer a.lock.Unlock()

	return n, success
}

func (a *Account) Deposit2(amount int64) (int64, bool) {
	a.lock.Lock()

	n, success := int64(0), false

	if !a.closed {
		new := a.balance + amount
		if new >= 0 {
			a.balance = new
			n, success = a.balance, true
		}

	}

	defer a.lock.Unlock()

	return n, success
}

func (a *Account) Close2() (int64, bool) {
	a.lock.Lock()

	n, success := int64(0), false

	if !a.closed {
		a.closed = true
		n = a.balance
		a.balance = 0
		success = true
	}

	defer a.lock.Unlock()

	return n, success
}
