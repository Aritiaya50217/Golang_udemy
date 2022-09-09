package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/new", newHandler)
	http.HandleFunc("/todos/create", createHandler)
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html", "todo.html"))

	// temp.Execute(w, []string{"learn golang", "practise exercise", "make coffice"})
	temp.Execute(w, readLines("todos.txt"))
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html", "new.html"))
	temp.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// log.Println(r)
	// log.Println(r.PostForm["item"])
	// log.Println(r.FormValue("item"))

	// [directory/file][111][000][000] => 0777
	// Permission
	file, err := os.OpenFile("todos.txt", os.O_APPEND, os.ModeAppend)
	check(err)
	_, err = fmt.Fprintln(file, r.FormValue("item"))
	check(err)
	// หลังจาก กรอกและ submit จะไปยัง /todos
	http.Redirect(w, r, "/todos", http.StatusFound)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(word string) []string {
	file, err := os.Open(word)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	// open file txt
	scanner := bufio.NewScanner(file)

	var lines []string
	// loop เพื่อดึงข้อมูล จาก slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}
