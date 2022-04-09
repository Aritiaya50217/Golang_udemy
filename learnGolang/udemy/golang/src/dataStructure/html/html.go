package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type TODO struct {
	UserId   int64  `json:"user_id"`
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

func getHtml() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(resp.Body)
	todos := []TODO{}
	todoDecoder.Decode(&todos)
	resp.Body.Close()
	todoEncoder := json.NewEncoder(os.Stdout)
	todoEncoder.Encode(todos)

	// index.html
	indexTemplateBytes, err := ioutil.ReadFile("/Users/MacbookPro/go/src/learnGolang/udemy/golang/src/dataStructure/html/index.html")
	if err != nil {
		return
	}
	// template
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)
}

func getHtmlFromFileJson() {
	resp, err := ioutil.ReadFile("/Users/MacbookPro/go/src/learnGolang/udemy/golang/src/dataStructure/html/todo.json")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(bytes.NewReader(resp))
	todos := []TODO{}
	todoDecoder.Decode(&todos)

	indexTemplateBytes, err := ioutil.ReadFile("/Users/MacbookPro/go/src/learnGolang/udemy/golang/src/dataStructure/html/index.html")
	if err != nil {
		return
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)
}

func main() {
	//getHtml()
	getHtmlFromFileJson()
}
