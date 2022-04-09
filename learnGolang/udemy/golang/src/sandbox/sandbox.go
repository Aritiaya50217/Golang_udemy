package main

import (
	"fmt"
)
//import "prime"

func pointer(){
	x := 70
	fmt.Println("Address x:",&x)

	var p *int
	// ค่าของ x (70)
	fmt.Println("&p :",&p)

	p = &x
	fmt.Println("p :" ,p)
	fmt.Println("value of *p : ",*p)

}

func returnPointer() *int{
	x := 4
	return &x
}

func receivePointer(r *int){
	*r = *r * 2
}

func newIntPointer() *int{
	var x int
	return &x
}

type KG float64
type LB float64

func declaration(){
	k := KG(3)
	//b := LB(3)
	result := k + 3
	fmt.Println(result)
}

func (lb LB) toKg() KG {
	return  KG(lb / 2.2046226218)
}
func declaredType() {
	k := KG(3)
	lb := LB(3)
	fmt.Println( k > lb.toKg())
}

func main() {
	// var
	var x int = 4
	fmt.Println(x)
	fmt.Printf("%T\n",x)

	var s []string
	fmt.Println(s)

	y := "xxx"
	fmt.Println(y)

	pointer()
	f := returnPointer()
	fmt.Print("Address x :",f , "\nvalue x : ",*f,"\n")

	r := 2
	receivePointer(&r)
	fmt.Println("result :",r)


	fmt.Println(newIntPointer() == newIntPointer()) // false เพราะ address จะเปลี่ยนไปเรื่อย ๆ
	// คล้ายกับ Built in package ของ Go
	fmt.Println(new(int) == new(int))

	declaration()

	declaredType()

	// package and initialization f  unction
	//for i := 0 ; i <= 100; i++ {
	//	x := rand.Intn(100000)
	//	fmt.Printf("%d , %t\n" ,x , prime.IsPrime(2))
	//}
}
