package main

import (
	"fmt"
	"reflect"
)

func array(){
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"manage",
	}
	fmt.Println("Length :",len(fruits) , fruits)
}
func gotCha(){
	twoSlots := [2]int{}
	threeSlots := [3]int{}
	fmt.Println("Two slots\t",reflect.TypeOf(twoSlots))
	fmt.Println("Three slots\t",reflect.TypeOf(threeSlots))
}

func ellipsis(){
	// การใช้ ...
	fruits := [...]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"manage",
	}
	fmt.Println("Length :",len(fruits) , fruits)
}

func indexInArray(){
	fruits := [...]string{
	3:	"apple",
	1:	"banana",
	0:	"papaya",
	4:	"orange",
	2:	"manage",
	}
	fmt.Println("Length :",len(fruits) , fruits[4],fruits[2])

	fmt.Println("true or false : ",[2]int{1,2} == [2]int{1,2})
	fmt.Println("true or false : ",[3]int{1,2,3} == [3]int{3,2,1})
}

func copyByValue(){
	fruits := [...]string{
		3:	"apple",
		1:	"banana",
		0:	"papaya",
		4:	"orange",
		2:	"manage",
	}
	dubFruits := fruits
	fmt.Println(fruits)
	fmt.Println("Before :",dubFruits)

	// edit value in array
	dubFruits[0] = "Watermelon"
	fmt.Println("After :",dubFruits)
}

func twoDimensional(){
	a := [2][2]int{{1,2},{3,4}}
	fmt.Println(a)
}

func main(){
	array()
	fmt.Println("--------------")
	gotCha()
	fmt.Println("--------------")
	ellipsis()
	fmt.Println("--------------")
	indexInArray()
	fmt.Println("--------------")
	copyByValue()
	fmt.Println("--------------")
	twoDimensional()
}