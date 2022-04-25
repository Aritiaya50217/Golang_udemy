package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byName []*Person
type customSort struct {
	Persons []*Person
	less    func(i, j *Person) bool
}

func (p customSort) Len() int {
	return len(p.Persons)
}

func (p customSort) Less(i, j int) bool {
	return p.less(p.Persons[i], p.Persons[j])
}

func (p customSort) Swap(i, j int) {
	p.Persons[i], p.Persons[j] = p.Persons[j], p.Persons[i]
}

// type Interface interface
// Len is the number of elements in the collection.
func (p byName) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p byName) Less(i, j int) bool {
	return p[i].name < p[j].name
}

// Swap swaps the elements with indexes i and j .
func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func printPerson(p []*Person) {
	for _, v := range p {
		fmt.Println(*v)
	}
}
func main() {
	p := []*Person{
		{"A", 22},
		{"B", 23},
		{"C", 24},
		{"D", 25},
		{"E", 22},
		{"F", 25},
		{"A", 21},
		{"B", 29},
	}
	printPerson(p)
	fmt.Println("-------------Before------------------")
	sort.Sort(byName(p))
	printPerson(p)

	fmt.Println("---------------Sorted-----------------")
	sort.Sort(customSort{Persons: p, less: func(i, j *Person) bool {
		if i.name != j.name {
			return i.name < j.name
		}
		if i.age != j.age {
			return i.age < j.age
		}
		return false
	}})
	printPerson(p)

	fmt.Println("-------------Reverse------------------")
	sort.Sort(sort.Reverse(customSort{Persons: p, less: func(i, j *Person) bool {
		if i.name != j.name {
			return i.name < j.name
		}
		if i.age != j.age {
			return i.age < j.age
		}
		return false
	}}))
	printPerson(p)
}
