package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func (m MyError) Error() string {
	return m.msg
}

func main() {
	var err error
	errors.New("sdf")
	fmt.Printf("%T, %#v", err, err)

	// check
	fmt.Println(errors.New("sdf") == errors.New("sdf"))

	err = MyError{"this is an error"}
	fmt.Printf("%T , %#v\n", err, err)

}
