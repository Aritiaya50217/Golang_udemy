package main

import (
	report "Reflection/report"
	"fmt"
	"reflect"
)

/*

	reflect คือ การ check type ของ value ขณะ runtime
	reflection จะส่งมอบข้อมูลพื้นฐาน ได้แก่ type , kind , value

	type ใช้ reflect.TypeOf เพื่อแสดงข้อมูล reflect.Type ขณะ runtime
	kind ใช้เมื่ออยากรู้ว่า original type คืออะไร
	value กรณีที่เราต้องการ reflect ข้อมูลใน interface{} ว่าแต่ละ field อยู่บนพื้นฐาน type อะไร
	จะใช้ reflect.ValueOf เพื่อเข้าถึง

*/

type MyStruct struct {
	a int
}

func testReflect() {
	look(MyStruct{a: 3})
}

func look(v interface{}) {
	fmt.Println("type :", reflect.TypeOf(v))
	fmt.Println("kind :", reflect.TypeOf(v).Kind())
	fmt.Println("value :", reflect.ValueOf(v))
}

type Person struct {
	FirstName string `report:"ชื่อ,uppercase"`
	Age       int    `report:"อายุ"`
}

type Employee struct {
	Name string
	Age  int
}

func main() {
	testReflect()

	// report
	fmt.Println(report.Text(Person{FirstName: "David", Age: 12}))
	fmt.Println(report.Text(Employee{Name: "Emily", Age: 11}))

	fmt.Println(report.Text(struct {
		Name string
		Age  int
	}{Name: "Emily", Age: 11}))

	// modified struct
	fmt.Println("===== Modified Struct =====")
	person := Person{FirstName: "David", Age: 12}
	fmt.Println("Before :", person)

	v := reflect.ValueOf(&person)
	fmt.Println(v) 
	// Field(n) => index field in person struct
	v.Elem().Field(0).SetString("DAVID")
	v.Elem().Field(1).SetInt(v.Elem().Field(1).Int() * 2) // Age : 12 => 24
	fmt.Println("After :", person)
	
}
