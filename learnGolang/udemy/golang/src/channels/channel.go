package main

import (
	"fmt"
	"math/rand"
	"time"
)

func training(ch chan bool) {
	rand.Seed(int64(time.Now().Nanosecond()))
	x := rand.Intn(1000)
	fmt.Println("training for :", x, "millisecond")
	time.Sleep(time.Duration(x) * time.Millisecond)
	ch <- true
}

func main() {
	// make channel
	done := make(chan bool)
	go training(done)
	trainingStatus := <-done
	fmt.Println("Done Status :", trainingStatus)

}
