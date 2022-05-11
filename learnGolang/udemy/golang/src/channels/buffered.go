package main

import (
	"fmt"
	"time"
)

// send string to channel
func pushLetter(ch chan<- string) {
	for i := "A"; i <= "G"; {
		fmt.Println("->", i)
		// send i to ch
		ch <- i
		i = string([]byte(i)[0] + 1)
	}
	fmt.Println("Close ch")
	close(ch)
}

func printLetter(ch chan string, done chan bool) {
	time.Sleep(10 * time.Millisecond)
	for v := range ch {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("\t| ", v)
	}
	done <- true
}
func main() {
	// channels มีขนาดเท่ากับ 3
	ch := make(chan string, 3)
	done := make(chan bool)

	// go routine
	go pushLetter(ch)
	go printLetter(ch, done)

	<-done
	fmt.Println("Done")
}
