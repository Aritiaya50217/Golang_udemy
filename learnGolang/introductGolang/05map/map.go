package main

import "fmt"

var product = make(map[string]float64)
func main(){
	fmt.Println("Product :",product)
	//add
	product["Macbook"] = 40000
	product["iPhone"] = 40001
	product["iPad"] = 40002
	fmt.Println(product)

	//del map
	delete(product, "iPad")
	fmt.Println("Product : ",product)

	//update
	product["iPhone"] = 20000
	fmt.Println("Product : ",product)

	// เข้าถึงข้อมูล
	value1 := product["iPhone"]
	fmt.Println("value1",value1)

	courseName := map[string]string{"101" : "Java" , "102":"Python"}
	fmt.Println("courseName",courseName)

}
