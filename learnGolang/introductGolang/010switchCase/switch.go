package main

import "fmt"


func color(x string)string{

	switch x {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("invalid value")
	}
	return x
}

func main(){
	input := 4
	switch input {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("invalid value")
	}

	c := color("blue")
	fmt.Println("color :",c)
}
