package main

import (
	"fmt"
	"log"
	"time"
)

func createF() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func createFList(list []int) []func() {
	result := []func(){}

	//result = append(result, func() {
	//	fmt.Println(list[0])
	//})
	//
	//result = append(result, func() {
	//	fmt.Println(list[1])
	//})

	for i, v := range list {
		i := i
		captureI := v
		result = append(result, func() {
			fmt.Printf("index : %v inside createF : %v\n", i, captureI)
		})
	}
	return result
}

func createFib() func(int) []int {
	fList := []int{0, 1, 1, 2, 3, 5}
	return func(n int) []int {
		if n > len(fList) {
			for n > len(fList) {
				// last index (n-1)
				lastIndex := len(fList) - 1
				fList = append(fList, fList[lastIndex]+fList[lastIndex-1])
			}
		}
		return fList[:n]
	}
}

func profileTime(f func(int) []int) func(int) []int {
	return func(a int) []int {
		start := time.Now()
		result := f(a)
		log.Printf("call with %d tooks %d ns", a, time.Now().Sub(start))
		return result
	}
}

func higherOrder() {
	fib := createFib()
	fmt.Println("-----------------")

	log.SetFlags(0)

	fib = profileTime(fib)

	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(60))
}

func anonymous() {
	// [f1(),f2()]
	fList := []func(){}
	fList = createFList([]int{10, 20, 30, 40, 50})
	for _, v := range fList {
		v()
	}
}

func main() {
	f := createF()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("-----------------------")
	anonymous()
	higherOrder()

}
