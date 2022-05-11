package main

import (
	"fmt"
	"sync"
)

var chDeposit = make(chan int)
var chBalance = make(chan int)

func bankSystem() {
	balance := 0
	for {
		if len(chDeposit) == 0 {
			select {
			case x := <-chDeposit:
				balance += x
			case chBalance <- balance:

			}
		} else {
			select {
			case x := <-chDeposit:
				balance += x
			}
		}
	}
}

func deposit(x int, w *sync.WaitGroup) {
	chDeposit <- x
	w.Done()
}

func finalBalance() int {
	return <-chBalance
}

func main() {
	go bankSystem()
	fmt.Println("Balance before deposit : ", finalBalance())
	w := &sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		w.Add(1)
		go deposit(100, w)
	}
	fmt.Println("Balance before waitgroup : ", finalBalance())
	w.Wait()

	fmt.Println("Balence :", finalBalance())
}
