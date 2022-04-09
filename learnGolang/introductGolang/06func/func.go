package main

import "fmt"

func Hello(){
	fmt.Println("Hello World")
}

func plus(v1,v2 int64)int64{
	return v1 + v2
}

func main(){
	Hello()

	res := plus(10,11)
	fmt.Println(res)
}