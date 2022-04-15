package main

import "fmt"

type Person struct {
	Name    string
	Surname string
}

func (p *Person) FullName() string {
	return p.Name + " " + p.Surname
}

type Employee struct {
	Person
	Id     string
	Office string
}

func (e *Employee) Detail() string {
	return "ID : " + e.Id + ". Office :" + e.Office + ".Fullname : " + e.FullName()
}

func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Programmer struct {
	Employee
	Language []string
}

func (p *Programmer) Detail() string {
	return "Programmer : " + p.Employee.Detail()
}

type Tester struct {
	Employee
	Tools []string
}

func (t *Tester) Detail() string {
	return "Tester : " + t.Employee.Detail()
}

func main() {
	david := Person{
		Name:    "David",
		Surname: "xxxx",
	}
	empDavid := Employee{
		Person: david,
		Id:     "123",
		Office: "Thailand",
	}
	progDavid := Programmer{
		Employee: empDavid,
		Language: []string{"golang", "Java", "Python"},
	}
	fmt.Printf("%+v", progDavid)

	oliver := Person{
		Name:    "Oliver",
		Surname: "xxx",
	}

	empOliver := Employee{
		Person: oliver,
		Id:     "456",
		Office: "Thailand",
	}
	testerOliver := Tester{
		Employee: empOliver,
		Tools:    []string{"Robot"},
	}
	fmt.Printf("%+v", testerOliver)
	fmt.Println(empDavid.IsSameOffice(&empOliver))

	fmt.Println(empDavid.IsSameOffice(&testerOliver.Employee))

	devidDetail := progDavid.Detail()
	fmt.Println("davidDetail :", devidDetail)
}
