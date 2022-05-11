package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Instructor string `json:"instructor"`
}

var courseList []Course

func init() {
	courseJson := `[
		{
			"id":1,
			"name" : "Python",
			"price": 2590,
			"instructor":"BorntoDev"
		},
		{
			"id":2,
			"name" : "JavaScript",
			"price": 0,
			"instructor":"BorntoDev"
		},
		{
			"id":3,
			"name" : "SQL",
			"price": 0,
			"instructor":"BorntoDev"
		}
	]`
	err := json.Unmarshal([]byte(courseJson), &courseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextId() int {
	highestID := -1
	for _, course := range courseList {
		if highestID < course.Id {
			highestID = course.Id
		}
	}
	return highestID + 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(courseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newCourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.Id = getNextId()
		courseList = append(courseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
