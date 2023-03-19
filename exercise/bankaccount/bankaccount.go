package account

import (
	"sync"
)

// Storage interface for Account struct type
type Storage interface {
	Close() (payout int64, ok bool)
	Balance() (balance int64, ok bool)
	Deposit(amount int64) (newBalance int64, ok bool)
}

// Account represent bank account that holds balance
type Account struct {
	balance int64
	mx      sync.RWMutex
	active  bool
}

// Open creates fresh account
func Open(initialDeposit int64) Storage {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, sync.RWMutex{}, true}
}

// Balance returns balance of account and true if account is active or 0 and false otherwise
func (b *Account) Balance() (balance int64, ok bool) {
	b.mx.RLock()
	defer b.mx.RUnlock()
	if !b.active {
		return
	}
	ok = true
	balance = b.balance
	return
}

// Deposit deposits given amount to account
func (b *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	b.mx.Lock()
	defer b.mx.Unlock()
	if !b.active {
		return
	}
	if b.balance+amount < 0 {
		return
	}
	b.balance += amount
	ok = true
	newBalance = b.balance
	return
}

// Close returns balance and sets Account active false
func (b *Account) Close() (payout int64, ok bool) {
	b.mx.Lock()
	defer b.mx.Unlock()
	if !b.active {
		return
	}
	payout += b.balance
	b.balance = 0
	b.active = false
	ok = true
	return
}
