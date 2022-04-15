package main

import (
	"flag"
	"fmt"
	"strconv"
)

var romanMap = map[string]int{
	"i":   1,
	"ii":  2,
	"iii": 3,
	"iv":  4,
	"v":   5,
}

type RomanAge struct {
	age string
}

func (r RomanAge) String() string {
	return strconv.Itoa(romanMap[r.age])
}

func (r *RomanAge) Set(value string) error {
	r.age = value
	return nil
}

func flagRomanAge() *RomanAge {
	romanAge := RomanAge{}
	flag.Var(&romanAge, "age", "help message for flagname")
	return &romanAge
}

var name = flag.String("name", "anonymous", "your name")
var romanAge = RomanAge{}
var age = flagRomanAge()

func main() {
	fmt.Println(*name)
	fmt.Println(age)
}
