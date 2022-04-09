package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	value string
	nodes []Node
}

func add(a, b float64) (c float64) {
	c = (a + b) / 2
	return c
}

func recursive(c int64) {
	fmt.Println(c)
	if c == 0 {
		return
	}
	recursive(c - 1)
	return
}

func leafNode() {
	p := Node{value: "p"}
	g := Node{value: "g"}
	b := Node{value: "b", nodes: []Node{p, g}}
	q := Node{value: "q"}
	s := Node{value: "s"}
	k := Node{value: "k"}
	r := Node{value: "r", nodes: []Node{q}}
	a := Node{value: "a", nodes: []Node{r, s}}
	root := Node{value: "a", nodes: []Node{b, a, k}}
	fmt.Println(root)
	result := outline([]string{}, &root)
	fmt.Println(result)
}

func outline(stack []string, n *Node) [][]string {
	stack = append(stack, n.value)
	result := [][]string{}
	if len(n.nodes) == 0 {
		result = append(result, stack)
	}
	for _, v := range n.nodes {
		result = append(result, outline(stack, &v)...)
	}
	return result
}

func main() {
	a := add(2, 4)
	fmt.Println("(a + b) / 2 =", a)

	// check type
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("---------------------")
	recursive(5)

	fmt.Println("---------------------")
	leafNode()

}
