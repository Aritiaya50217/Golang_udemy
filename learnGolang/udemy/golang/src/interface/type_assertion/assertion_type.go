package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Write(p []byte) (n int , err error)
	var w io.Writer
	w = os.Stdout
	//interface.(type / interface)
	result := w.(*os.File)
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T , %#v\n", result, result)

	result2 := w.(*bytes.Buffer)
	fmt.Printf("%T , %#v\n", w, w)
	fmt.Printf("%T , %#v\n", result2, result2)
}
