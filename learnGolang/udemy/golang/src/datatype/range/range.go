package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func rangeLoop(){
	x := "ทดสอบ"
	for i , v := range x {
		fmt.Printf("%d, %c\n",i,v)
	}
	fmt.Println("utf8.RuneCountInString :",utf8.RuneCountInString(x))

	for i:= 0 ; i < len(x);{
		r , s := utf8.DecodeRuneInString(x[i:])
		i += s
		fmt.Printf("%c-",r)
	}
	// check "สอ" ว่าอยู่ใน x ไหม
	finder := "สอ"
	fmt.Println()
	//แบบที่ 1
	fmt.Println(bytes.Count([]byte(x),[]byte(finder)))
	//แบบที่ 2
	fmt.Println(strings.Count(x ,finder))

	buff := strings.Builder{}

	buff.WriteRune('h')
	buff.WriteRune('i')
	fmt.Println(buff.String())
	// ASCII to Int
	atoi,_ := strconv.Atoi("123")
	fmt.Println(atoi,reflect.TypeOf(atoi))

	// Int to ASCII
	itoa := strconv.Itoa(123)
	fmt.Println(itoa,reflect.TypeOf(itoa))

}

func main(){
	rangeLoop()
}
