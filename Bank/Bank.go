package Bank

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Account struct {
	balance          int
	mu               sync.Mutex
	transactionCount uint64
}

func (acc *Account) Deposit(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += amount
	atomic.AddUint64(&acc.transactionCount, 1)
}

func (acc *Account) Withdraw(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.balance >= amount {
		acc.balance -= amount
		atomic.AddUint64(&acc.transactionCount, 1)
	}
}

func (acc *Account) GetBalance() int {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.balance
}

func Demo() {
	var account Account
	account.Deposit(100)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			account.Deposit(100)
		}()
		go func() {
			defer wg.Done()
			account.Withdraw(50)
		}()
	}

	wg.Wait()

	fmt.Println("Final balance:", account.GetBalance())
	fmt.Println("Total transactions:", atomic.LoadUint64(&account.transactionCount))
}
