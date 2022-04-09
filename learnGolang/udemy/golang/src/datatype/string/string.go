package main

import "fmt"

func lengthStr(){
	eng := "Hello"
	th := "สวัสดี"

	fmt.Println(eng,len(eng) , string(eng[0]))
	fmt.Println(th,len(th),string(th[0]))
}
func stringRepresentation(){
	x := "Hi-สวัสดี" // []byte{0,1...20}
	y := []byte{0x48,0x69,0x2d,0xe0,0xb8,0xaa,0xe0,0xb8,0xa7,0xe0,0xb8,0xb1,0xe0,0xb8,0xaa,0xe0,0xb8,0x94,0xe0,0xb8,0xb5}
	fmt.Println(x , len(x))
	fmt.Println(string(y))
	//fmt.Println(string(0x41))
	//for i := 0; i < len(x); i++ {
	//	// %x => ฐาน 16
	//	fmt.Printf("%d = 0x%x\n",i,x[i])
	//}
}

func uniCode8(){
	x := "Hi-สวัสดี" // []byte{0,1...20}
	y := []byte{0x48,0x69,0x2d,0xe0,0xb8,0xaa,0xe0,0xb8,0xa7,0xe0,0xb8,0xb1,0xe0,0xb8,0xaa,0xe0,0xb8,0x94,0xe0,0xb8,0xb5}
	fmt.Println(x , len(x))
	fmt.Println(string(y))
	// ตัวอย่างอักษรภาษาไทยมีขนาด 3 byte(รู้หลังจากการ encode) จึงต้องใส่ระยะห่างไป 3
	fmt.Println(string(y[3]),string(y[3:6]))  // ส
	fmt.Println("utf-8 :","\xe0\xb8\xaa") // ส

	// 1110 0010 1010 (หา byte ของ ส) (unicode) => utf-8
	// 11100000 10111000 10101010 = utf-8
	// 111000001011100010101010 = e0b8aa (เลขฐาน 16)
}

func exRune(){
	x := "Hi-สวัสดี" // []byte{0,1...20}
	y := []byte{0x48,0x69,0x2d,0xe0,0xb8,0xaa,0xe0,0xb8,0xa7,0xe0,0xb8,0xb1,0xe0,0xb8,0xaa,0xe0,0xb8,0x94,0xe0,0xb8,0xb5}
	z := []rune(x)
	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Printf("%q",z[5])  // ั
}

func conversions(){
	// Case 1 : Integer to String
	ex1 := string(0x265e)
	fmt.Println("Case 1 : Integer to String")
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%x ",ex1[i])
	}
	fmt.Println("\n\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)

	// Case 2 : slice to byte to string
	fmt.Println("\nCase 2 : []byte to string")
	ex2 := []byte{0xe2 , 0x99 , 0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String)

	// Case 3 : slice of rune to string
	fmt.Println("\nCase 3: []rune to string")
	ex3 := []rune{0x2654,0x2655,0x2655,0x2657,0x2658,0x2659}
	fmt.Println(string(ex3))

	// Case 4 : string  to slice of byte
	fmt.Println("\nCase 4: string to []byte")
	ex4 := []byte("Hello ♔")
	fmt.Println(ex4)

	// Case 5 : string to a slice  of runes
	fmt.Println("\nCase 5 : string to []rune ")
	ex5 := []rune("hello")
	fmt.Println(ex5)
}

func main(){
	lengthStr()
	stringRepresentation()
	fmt.Println("---------------------")
	uniCode8()
	fmt.Println("---------------------")
	exRune()
	fmt.Println("---------------------")
	conversions()
}