package main

import (
	"fmt"
	"os"
)

func ExPanic() {
	x := []int{1, 2, 3}
	fmt.Println(x[1])
	//panic("for no reason")
	f, err := os.Open("sasasas")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}

var cygwinPath = os.Getenv("CYGWIN")

func init() {
	fmt.Println("init cygwing path :", cygwinPath)
	if cygwinPath == "" {
		panic("cannot set cygwin path. Make sure you have one")
	}
}

func cygwing() {
	// utilities for cygwing
	// path cygwin ?
	fmt.Println("Hello to cygwin utilities")

}

func main() {
	ExPanic()
	cygwing()
}
