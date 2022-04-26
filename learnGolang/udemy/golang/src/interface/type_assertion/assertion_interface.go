package main

import (
	"fmt"
	"io"
)

type MyIo struct {
}

func (m MyIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

//
//func (m MyIo) Read(p []byte) (n int, err error) {
//	return len(p), nil
//}

func Value(key interface{}) {
	if keyAsString, ok := key.(string); ok {
		fmt.Println("assert ok value is :", keyAsString)
		return
	}
	fmt.Println("assert not ok")
}

func main() {
	var w io.Writer

	//w = os.Stdout
	//result := w.(http.Handler)
	//fmt.Printf("%T , %#v\n", w, w)
	//fmt.Printf("%T ,%#v\n", result, result)

	// panic
	w = MyIo{}
	result, ok := w.(io.Reader)
	fmt.Printf("%T , %#v\n", w, w)
	fmt.Printf("%T ,%#v,%#v\n", result, result, ok)

	//gin framework
	fmt.Println("--------- gin --------")
	Value(MyIo{})

}
