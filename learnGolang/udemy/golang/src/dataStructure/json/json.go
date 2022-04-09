package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"` // omitempty กรณีไม่มีการใส่ข้อมูล จะไม่แสดงคีย์
	Completed bool   `json:"completed"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Users []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  string `json:"company"`
}

var data = `[{
		"userId":1,
		"id":	 1,
		"title":	"delectus aut autem",
		"completed": false
	},
	{    "userId":2,
	     "id":	 2,
	     "title":	"delectus aut autem",
	     "completed": false
	},
	{    "userId":2,
	     "id":	 3,
	     "completed": false
	}
]`

func jsonUnmarshal() {
	datastruct := []Todo{}
	v := &datastruct
	fmt.Println(datastruct)

	json.Unmarshal([]byte(data), v)
	fmt.Println(datastruct)

}

func jsonMarshal() {
	datastruct := Users{}
	v := &datastruct
	fmt.Println(datastruct)

	json.Unmarshal([]byte(data), v)
	fmt.Println(datastruct)

	// dataStruct --> std.Output
	result, err := json.MarshalIndent(datastruct, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))
}

func httpGet() {
	// https://jsonplaceholder.typicode.com/todos
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	dataStruct := Users{}
	v := &dataStruct
	json.Unmarshal(body, v)
	fmt.Println(len(dataStruct))
	dataStruct[0].Name = "Oil"

	// dataStruct --> std.Output
	result, err := json.MarshalIndent(dataStruct, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))
}
func jsonEncoder() {
	// https://jsonplaceholder.typicode.com/todos
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}

	// json.Decode --> dataStruct
	jsonDecoder := json.NewDecoder(resp.Body)
	dataStruct := Users{}
	jsonDecoder.Decode(&dataStruct)
	resp.Body.Close()
	fmt.Println(len(dataStruct))
	dataStruct[0].Name = "Oil"

	// dataStruct --> std.Output
	//result, err := json.MarshalIndent(dataStruct, "", "    ")
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(result))

	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(dataStruct)
}

func main() {
	jsonUnmarshal()
	jsonMarshal()
	httpGet()
	httpGet()
	jsonEncoder()
}
