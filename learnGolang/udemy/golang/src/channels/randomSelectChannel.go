package main

import (
	"fmt"
	"time"
)

func running(name string, ch <-chan int) {
	for {
		fmt.Println(name, <-ch)
	}

}
func main() {

	// tricker1 := time.NewTicker(1 * time.Millisecond)
	// tricker2 := time.NewTicker(2 * time.Millisecond)
	// tricker3 := time.NewTicker(3 * time.Millisecond)
	// tricker4 := time.NewTicker(4 * time.Millisecond)

	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case v := <-tricker1.C:
	// 		fmt.Println("ch1 :", v.Nanosecond())
	// 	case v := <-tricker2.C:
	// 		fmt.Println("ch2 :", v.Nanosecond())
	// 	case v := <-tricker3.C:
	// 		fmt.Println("ch3 :", v.Nanosecond())
	// 	case v := <-tricker4.C:
	// 		fmt.Println("ch4 :", v.Nanosecond())

	// 	}
	// 	time.Sleep(500 * time.Microsecond)

	// }
	// time.Sleep(200 * time.Millisecond)

	after2Second := time.After(2 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case v := <-after2Second:

			fmt.Println("Got value form after2Second channel", v.Second())

		default:
			fmt.Println("Default")
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Done")
}
