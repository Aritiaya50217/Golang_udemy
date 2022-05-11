package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// json.Marshal --> object to json

type Employee struct {
	Id    int
	Name  string
	Phone string
	Email string
}

// json.Unmarshal --> json to object
type Employee2 struct {
	Id    int
	Name  string
	Phone string
	Email string
}

func main() {
	data, _ := json.Marshal(&Employee{
		101,
		"AAA",
		"090909090",
		"user1@gmail.com",
	})
	fmt.Println("-----Marshal-----")
	fmt.Println(string(data))

	e := Employee2{}
	err := json.Unmarshal([]byte(`{"Id": 101,"Name":"BBB","Phone":"09990909","Email":"user2@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----Unmarshal-----")
	fmt.Println("Name :", e.Name)
}
