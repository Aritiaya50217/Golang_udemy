package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      = sync.RWMutex{}
	balance = 100
)

//  51137356 --> no mutex
// 510639090 --> use mutex
//  51279780 --> use RWMutex

func readX(wg *sync.WaitGroup) int {
	mu.RLock()
	defer func() {
		mu.RUnlock()
		wg.Done()
	}()
	time.Sleep(50 * time.Millisecond)
	return balance
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("took :", time.Since(start).Nanoseconds())
	}()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readX(wg)
	}
	wg.Wait()
}
