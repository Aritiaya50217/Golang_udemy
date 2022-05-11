package main

import "fmt"

func sender2(ch chan<- int) {
	// send int to channel
	ch <- 88

}

func f(ch chan int) {
	ch <- 99
	close(ch)
}

func receiver2(ch <-chan int, done chan bool) {
	fmt.Println("Receive ", <-ch)
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	// gorountines
	go sender2(ch)
	go receiver2(ch, done)

	<-done
	fmt.Println("Done")
}
