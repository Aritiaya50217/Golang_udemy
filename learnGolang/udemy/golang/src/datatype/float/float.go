package main

import (
	"fmt"
	"math/cmplx"
)


func floatNumber(){
	var x float64
	fmt.Println(x,-x , 1/x , 1/-x , x/x)
}
func complexNumber(){
	x := complex128(complex(7,3))
	var y complex128 = complex(7,3)
	fmt.Println("x :",x)
	fmt.Println("y :",y)
	fmt.Println("x + y :",x + y)
	fmt.Println("x * y :",x * y)
	// พีทาโกรัท
	fmt.Println(cmplx.Abs(complex(3,4)))
}

func main(){
	floatNumber()
	complexNumber()
}