package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func findId(id int) (*Course, int) {
	for i, course := range courseList {
		if course.Id == id {
			return &course, i
		}
	}
	return nil, 0
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/")
	id, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, listItemIndex := findId(id)
	if course == nil {
		http.Error(w, fmt.Sprintf("no course with id %d", id), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		courseJson, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJson)

	case http.MethodPut:
		var updatedCourse Course
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byteBody, &updatedCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedCourse.Id != id {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updatedCourse
		courseList[listItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/course/", courseHandler)
	http.HandleFunc("/course", coursesHandler)
	http.ListenAndServe(":5000", nil)
}
