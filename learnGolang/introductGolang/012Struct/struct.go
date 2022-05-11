package main

import "fmt"

type Employee struct {
	Id    int64
	Name  string
	Phone string
}

func main() {
	employee1 := Employee{
		01,
		"AAA",
		"0123456788",
	}

	fmt.Println("employee1 :", employee1)

	fmt.Println("-------------------------------")

	employeeList := [3]Employee{}
	employeeList[0] = Employee{
		01,
		"BBB",
		"09090909000",
	}
	employeeList[1] = Employee{
		02,
		"CCC",
		"09090909000",
	}
	employeeList[2] = Employee{
		03,
		"DDD",
		"09090909000",
	}
	fmt.Println("empList :", employeeList)
}
