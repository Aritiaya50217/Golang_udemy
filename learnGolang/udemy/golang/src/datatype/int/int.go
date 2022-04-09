package main

import (
	"fmt"
	"math"
	"strconv"
)

// integers ---> signed  (8,16,32,62)
//     "    ---> unsigned (8,16,32,64)

func checkMin(){
	fmt.Println("Min :",math.MinInt8)
}

func checkMax(){
	fmt.Println("Max :",math.MaxInt8)
}

func checkUint(){
	// MinInt8 + MaxInt8 = Uint8
	fmt.Println("Uint8 :",math.MaxUint8)
}

func secondCompliment(){
	x := -5
	fmt.Printf("%b\n",x)
	fmt.Printf(strconv.FormatInt(int64(x), 2))

	// หาค่า -5 กรณีเป็นเลขฐานสอง
	fmt.Println(bitInt(int8(x)))
}

// 1111 1111
// 0000 0001

func bitInt( x int8) [8]byte{
	 var result [8]byte // 1111 1011
     for i := 0 ; i < len(result); i++ {
		 result[7-i] = byte(x & 1)
		 x = x >> 1
	 }
	 return result
}

func integerOverFlow(){
	x := 127
	y := 2
	fmt.Println(int8(x+y)) // 129
	fmt.Printf("%b\n",128)
	fmt.Printf("%b",int8(x+y))
}

func main(){
	checkMin()
	checkMax()
	checkUint()
	secondCompliment()
	integerOverFlow()
}