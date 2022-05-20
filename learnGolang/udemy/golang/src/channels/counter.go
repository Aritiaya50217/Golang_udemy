package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const N = 1000
	add := 0           // N
	sub := -0          // -N
	mu := sync.Mutex{} // Mutual Exclusion
	for i := 0; i < N; i++ {
		add++
		go func() {
			defer func() {
				// lock
				mu.Lock()
				sub-- // critical section
				// unlock
				mu.Unlock()
			}()
		}()
	}
	for {
		h, m, s := time.Now().Clock()
		fmt.Printf("%d:%d:%d : add %d , sub %d\r", h, m, s, add, sub)
		if add == N && sub == -N {
			fmt.Println("Success")
			break
		}
	}
	fmt.Println("main done", add, sub)
}
