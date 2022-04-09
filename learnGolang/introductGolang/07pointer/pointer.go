package main

import "fmt"

func zeroValue(i int) {
	i = 0

}

func zeroPointer(pointer *int){
	*pointer = 0
}

func main(){
	i := 1
	fmt.Println(i)

	zeroValue(i)
	fmt.Println("Value : ",i)

	zeroPointer(&i)
	fmt.Println("Address : ",i)

}