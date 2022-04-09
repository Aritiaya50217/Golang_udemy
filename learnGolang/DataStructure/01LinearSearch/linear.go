package main

import "fmt"

func linearSearch(datalist []int , key int)bool{
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main(){
	items := []int{9,3,5,6,7,2,5,9,4,4,4}
	fmt.Println(linearSearch(items,44))
}