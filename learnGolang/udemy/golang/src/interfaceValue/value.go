package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var w io.Writer

func printBuff(w io.Writer) {
	if w != nil {
		fmt.Println(w.Write([]byte("hello")))
	}
}

func main() {
	fmt.Printf("type : %T\nvalue : %v\n", w, w)

	w = os.Stdout
	fmt.Println("---- os.Stdout ---- ")
	fmt.Printf("type : %T\nvalue : %v\n", w, w)

	w = &bytes.Buffer{}
	fmt.Println("---- bytes.Buffer ---- ")
	fmt.Printf("type : %T\nvalue : %v\n", w, w)

	var byteBuff io.Writer
	printBuff(byteBuff)

}
