package main

import "fmt"

func forLoop(){
	// for initialization; condition; post {
	//     body
	// }
	for i := 1; i < 5; i++{
		fmt.Println(i)
		if i == 3 {
			fmt.Println("I am about to terminate")
		}
	}
}

func main(){
	forLoop()
}
