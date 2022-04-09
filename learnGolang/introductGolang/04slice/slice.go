package main

import "fmt"

func main(){
	var courseName []string
	courseName = []string{"Java","Python"}
	fmt.Println(courseName)
	courseName = append(courseName,"C")
	fmt.Println(courseName)


	courseWeb := courseName[0:1]
	fmt.Println(courseWeb)
}
