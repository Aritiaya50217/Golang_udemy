package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("016file/test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line : %d --> %s\n", count, line)
		count++
	}

	// write
	data1 := []byte("hello gopher")
	err = os.WriteFile("016file/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	create, err := os.Create("016file/data2.txt")
	if err != nil {
		panic(err)
	}

	defer create.Close()

	data2 := []byte("Golang")
	os.WriteFile("016file/data2.txt", data2, 0644)

}
