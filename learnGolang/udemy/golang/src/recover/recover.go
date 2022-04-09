package main

import "fmt"

func A(x int) {
	defer func() {
		fmt.Println("defer in A")
	}()
	B(x)
	fmt.Println("Hello in A")
}

func B(x int) {
	defer func() {
		fmt.Println("defer in B")
	}()
	C(x)
}

func C(x int) {
	defer func() {
		fmt.Println("defer in C")
		if p := recover(); p != nil {
			fmt.Println("handling panic :", p)
		}
	}()
	if x == 4 {
		panic("for no reason")
	}
}

func main() {
	A(4)
}
