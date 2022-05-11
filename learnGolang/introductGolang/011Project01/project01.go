package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		msg, _ := fmt.Scanf("%v must number only :", promt)
		panic(msg)
	}
	return value
}

func getOperator() string {
	fmt.Println("operator is (+ , - , * , / )")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(x, y float64) float64 {
	return x + y
}
func sub(x, y float64) float64 {
	return x - y
}
func multi(x, y float64) float64 {
	return x * y
}
func division(x, y float64) float64 {
	return x / y
}

func main() {
	value1 := getInput(" value1 : ")
	value2 := getInput(" value2 : ")

	var result float64
	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = sub(value1, value2)
	case "*":
		result = multi(value1, value2)
	case "/":
		result = division(value1, value2)
	default:
		fmt.Println("invalid operator :")
	}
	fmt.Println("result :", result)
}
