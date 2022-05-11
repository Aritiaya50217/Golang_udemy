package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func process(c chan string, data string) {
	c <- data

}
func main() {
	fmt.Println("------go routine-----")
	go f("Hello Gopher")
	f("message")
	time.Sleep(5 * time.Second)

	fmt.Println("----Channel---")
	// make(chan ,type , ช่องที่จอง)
	ch := make(chan string, 1)
	go process(ch, "Data1")
	fmt.Println("Receive value :", <-ch)
}
