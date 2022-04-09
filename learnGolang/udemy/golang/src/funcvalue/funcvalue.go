package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}
func apply(a, b float64, op func(float64, float64) float64) (float64, error) {
	if op == nil {
		return math.NaN(), fmt.Errorf("apply : nil operation")
	}
	return op(a, b), nil
}
func main() {
	a, _ := apply(1, 2, add)
	b, _ := apply(1, 2, sub)
	c, _ := apply(1, 8, nil)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
