package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Name string
	Age  int64
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func templateSolution() {
	filicity := Person{"Filicity", 24}
	oliver := Person{"Oliver", 25}

	people := []Person{filicity, oliver}

	const greetPerson = `Hi I am {{.Name | upperCase}}. {{.Age}} years old {{"\n"}}`
	const greetPeople = `{{range .}} Hi I am {{.Name | upperCase}}. {{.Age}} years old {{"\n"}} {{end}} `

	greetTemplate := template.Must(template.New("greet From Person").Funcs(template.FuncMap{"upperCase": upperCase}).Parse(greetPerson))
	greetPeopleTemplate := template.Must(template.New("greet From People").Funcs(template.FuncMap{"upperCase": upperCase}).Parse(greetPeople))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)

	fmt.Println("----------------------------")
	greetPeopleTemplate.Execute(os.Stdout, people)
}
func main() {
	templateSolution()

}
