package main

import "fmt"

func createMap(){
	items := map[string]int{
		"pen":		4,
		"pencil":	5,
	}
	fmt.Println(items)
	fmt.Println("pencil :",items["pencil"])
}

func makeMap(){
	items := map[string]int{
		"pen":		4,
		"pencil":	5,
	}

	orders := make(map[string]int)
	orders["pen"] = 12
	orders["pencil"] = 1

	x := items
	x["ruler"] = 7
	fmt.Println("orders :",orders)
	fmt.Println("x :",x)

	delete(orders,"pencil")
	fmt.Println("After Delete pencil :",orders)
}

func checkValue(){
	items := map[string]int{
		"pen":		4,
		"pencil":	5,
	}

	fmt.Println("items :",items)
	v,ok := items["pen"]
	if !ok {
		fmt.Println("Not exist",v)
	}else {
		fmt.Println("value :",v)
	}

}

func main(){
	createMap()
	makeMap()
	checkValue()
}