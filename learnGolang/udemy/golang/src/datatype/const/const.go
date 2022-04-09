package main

import (
	"fmt"
	"reflect"
)

func constant(){
	var x = 2
	const y = 5 - 6
	fmt.Println(x / y)
}

func unTypeConst(){
	const persons = 4
	toffee := 5 / persons
	cost := 2.0 / persons
	fmt.Println(toffee , reflect.TypeOf(toffee))
	fmt.Println(cost , reflect.TypeOf(cost))

}

const (
	Sunday  	=	iota + 1 // constant generator
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


func main(){
	constant()
	unTypeConst()

	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday)
}