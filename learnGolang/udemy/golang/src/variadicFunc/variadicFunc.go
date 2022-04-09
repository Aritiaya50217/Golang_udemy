package main

import "fmt"

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	x := []interface{}{"cc", "bbb", "mmm", 23232}
	printEachLine(x...)
}
