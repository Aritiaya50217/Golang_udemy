package main

import "fmt"

func add(x, y float64) {
	result := x + y
	fmt.Println("result :", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i :", i)
	}
}

func deferLoop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j :", j)
	}
}

func main() {
	// defer ส่วนแสดงผลลำดับท้ายสุด
	//defer fmt.Println("End")
	//add(10.0, 29.99)

	loop()
	fmt.Println("-----defer----")
	deferLoop()

}
