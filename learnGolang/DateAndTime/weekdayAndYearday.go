package main

import (
	"fmt"
	"time"
)

func main() {
	// "y-m-d-t-s"
	t, _ := time.Parse("2006 01 02 15 04", "2015 11 11 16 50")
	fmt.Println(t.YearDay())
	fmt.Println(t.Weekday())

	t2, _ := time.Parse("2006 01 02 15 04", "2022 04 27 14 43")
	fmt.Println(t2.YearDay())
	fmt.Println("Day : ", t2.Weekday())

}
