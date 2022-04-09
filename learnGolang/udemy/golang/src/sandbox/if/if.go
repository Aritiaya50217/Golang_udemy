package main

import "fmt"

func ifCase(){
	x := 5

	if x < 5{
		fmt.Println("x greater than 5 .")
	}else {
		fmt.Println("x less than 5 .")
	}
}

func switchCase(){
	x := 1
	switch x {
	case 1 :
		fmt.Println("One")
		fallthrough // result One Two
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Empty")
	}
}

func  exSwitchCase(){
	x := 20
	switch true {
	case 1 <= x && x < 100 :
		fmt.Println("deci")
	case 100 <= x && x < 1000:
		fmt.Println("centi")
	case 1000 <= x:
		fmt.Println("meter")
	default:
		fmt.Println("I am default")
	}
}

func main(){
	ifCase()
	switchCase()
	exSwitchCase()
}
