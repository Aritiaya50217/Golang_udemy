package main

import (
	"fmt"
	"sync"
)

var (
	mu      = sync.Mutex{}
	balence = 0
)

func deposit(x int, w *sync.WaitGroup) {
	// start --> Lock
	// end --> Unlock
	mu.Lock()
	defer func() {
		mu.Unlock()
		w.Done()
	}()

	balence += x
	addBonus(1)

}

func AddBonus(x int) {
	mu.Lock()
	defer func() {
		mu.Unlock()

	}()
	addBonus(x)
}

func addBonus(x int) {
	balence += x
}

func finalBalance() int {
	mu.Lock()
	defer func() {
		mu.Unlock()

	}()
	return balence
}

func main() {
	w := &sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		w.Add(1)
		go deposit(100, w)
	}
	w.Wait()

	fmt.Println("Balence :", finalBalance())
}
