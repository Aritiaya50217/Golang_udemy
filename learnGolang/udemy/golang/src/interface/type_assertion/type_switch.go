package main

import "fmt"

func testValue(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("string value : ", v)
		return
	}

	if v, ok := x.(int); ok {
		fmt.Println("int value : ", v)
		return
	}

	if v, ok := x.(bool); ok {
		fmt.Println("boolean value : ", v)
		return
	}
	if v, ok := x.(float64); ok {
		fmt.Println("float value : ", v)
		return
	}
	fmt.Println("No match any ")
}

func testSwitchCase(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Println("string value: ", v)
	case int:
		fmt.Println("int value: ", v)
	case bool:
		fmt.Println("boolean value : ", v)
	case float32, float64:
		fmt.Println("float64 value :", v)
	default:
		fmt.Println("No match any.")
	}
}
func main() {
	testValue("sdf")
	testValue(123)
	testValue(true)
	testValue(1.1111)

	fmt.Println("\n------switch case----")
	testSwitchCase("oil")

}
