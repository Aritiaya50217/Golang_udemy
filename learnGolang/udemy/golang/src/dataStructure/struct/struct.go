package main

import (
	"fmt"
	"learnGolang/udemy/golang/src/dataStructure/hr"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func structure() {
	// แบบที่ 1
	//x := struct {
	//	name string
	//	age  int
	//}{"Lisa",23}

	// แบบที่ 2
	x := .Person{Name: "Lisa", Age: 24}
	fmt.Println(x)
	fmt.Println("Name :", x.Name)
}

func structAndPointer() {
	x := .Person{"Test", 20}
	fmt.Println("Before :", x)

	y := &x
	y.Name = "Oliver"
	fmt.Println("After :", x)
	fmt.Println(y)
}

func checkStruct() {
	x := .Person{"name1", 20}
	y := .Person{"name2", 20}

	fmt.Println("x == y :", x == y)
}

func embeddingStruct() {
	fillicity := .Employee{
		Person:      .Person{"Lisa", 25},
		Designation: "Programmer",
	}
	fmt.Printf("%+v\n", fillicity)
	fmt.Println(fillicity.Person.Name)
	fmt.Println(fillicity.Person.Age)
}

func structRecursive() {
	root := BinaryTree{value: 2}
	left := BinaryTree{value: 1}
	right := BinaryTree{value: 3}

	root.left = &left
	root.right = &right

	showDF(&root)
}

func showDF(node *BinaryTree) {
	if node != nil {
		showDF(node.left)
		fmt.Println(node.value)
		showDF(node.right)
	}
}
func main() {
	structure()
	structAndPointer()
	checkStruct()
	embeddingStruct()
	structRecursive()
}
