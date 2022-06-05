package example2

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	name    string
	lock    sync.Mutex
}

func (a *Account) Withdraw(wg *sync.WaitGroup, amount int) {
	defer wg.Done()
	a.lock.Lock()
	a.balance -= amount
	a.lock.Unlock()
}

func (a *Account) Deposit(wg *sync.WaitGroup, acmount int) {
	defer wg.Done()
	a.lock.Lock()
	a.balance += acmount
	a.lock.Unlock()
}

func (a *Account) GetBalance() int {
	return a.balance
}

func Example() {
	var wg sync.WaitGroup
	var account1 Account
	account1.name = "account1"

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go account1.Deposit(&wg, 100)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		account1.Withdraw(&wg, 100)
	}

	wg.Wait()
	fmt.Printf("Balance: %d\n", account1.GetBalance())
}
