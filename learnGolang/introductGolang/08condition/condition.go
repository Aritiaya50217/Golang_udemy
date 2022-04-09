package main

import "fmt"

func grade(x int64)string{

	if x >= 80 {
		return "A"
	}else if x >= 70 {
		return "B"
	}else if x >= 60 {
		return "C"
	}else if x >= 50 {
		return "D"
	}else {
		return "F"
	}
}

func main(){
	myMoney := 100
	if myMoney > 100 {
		fmt.Println("myMoney greater than 100")
	}else{
		fmt.Println("myMoney less than 100 ")
	}


	fmt.Println("Grade :",grade(98))
}