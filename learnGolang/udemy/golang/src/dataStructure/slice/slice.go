package main

import "fmt"

func slice(){
	fruits := [5]string{"apple","banana","papaya","orange","mango"}
	myFavor := fruits[1:4]
	fmt.Println("Length :",len(myFavor) ,"\n","cap :", cap(myFavor),"\n","slice :",myFavor)

	fmt.Println("---------------")
	yourFavor := myFavor
	fmt.Println("Length :",len(yourFavor) ,"\n","cap :", cap(yourFavor),"\n","slice :",yourFavor)

	fmt.Println("---------------")
	// change value in yourFavor
	yourFavor[0] = "guava"
	fmt.Println("Length :",len(yourFavor) ,"\n","cap :", cap(yourFavor),"\n","slice :",yourFavor)

}

func makeSlice(){
	// make([]type , len , cap)
	s := make([]int,5,10)
	s[0] = 6
	fmt.Println("Length : ",len(s))
	fmt.Println("cap : ",cap(s))
	fmt.Println("slice:",s)
}

func memoryAllocation(){
	s := make([]int,99)
	fmt.Println("Length slice :",len(s))
	fmt.Println("Cap :",cap(s))
	fmt.Println(s)
}

func appending(){
	x := [10]int{}
	a := x[0:5]
	b := x[5:7]
	for i := range a {
		a[i] = i * i
	}
	b[0] = 98
	b[1] = 99
	a = append(a,b...)
	a[len(a)-1] = 25
	a = append(a,70,72,73,74)
	fmt.Println("x :",x)
	fmt.Println("a :","len :",len(a),"cap :",cap(a),a)
	fmt.Println("b :","len :",len(b),"cap :",cap(b),b)


}

func main(){
	slice()
	makeSlice()
	memoryAllocation()
	appending()
}
