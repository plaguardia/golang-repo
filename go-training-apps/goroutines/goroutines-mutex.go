package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var balance int

func deposit(money int, w *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("depositing %d dollars to account with balance %d\n", money, balance)
	balance += money
	mutex.Unlock()
	w.Done()
}

func withdraw(money int, w *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("withdrawing %d dollars to account with balance %d\n", money, balance)
	balance -= money
	mutex.Unlock()
	w.Done()
}

func main() {

	//mutex is used to put a lock onto the memory being used by a function
	// real ex) you have multiple goroutines that are trying to append to the same array at the same time! (this is dangerous) you can lock the memory until individual processes are complete
	// "race conditions" - certain goroutines need to know about eachother
	balance = 1000

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go withdraw(700, wg)
	go deposit(500, wg)
	wg.Wait()

	fmt.Printf("Your account new balance is %d\n", balance)

}
