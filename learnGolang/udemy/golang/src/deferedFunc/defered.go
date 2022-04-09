package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func defered() {
	// defer แสดงผลหลังสุด
	fmt.Println("a")
	defer fmt.Println("last")
	fmt.Println("c")
}

func exdefer() {
	x := 1
	defer fmt.Println("a")
	fmt.Println(x)
}
func add() {
	x := 1
	addX := func(a int) int {
		x = x + a
		return x
	}
	defer fmt.Println(addX(x))
	fmt.Println(x)
}

func stopWatch(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("function %s : took tiem %d\n", name, time.Now().Sub(start).Milliseconds())
	}
}

func loadingImage() {
	defer stopWatch("loadingImage")()
	time.Sleep(3 * time.Millisecond)
	fmt.Println("loadingImage : done")
}

func GetHttp() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func main() {
	defered()
	exdefer()
	fmt.Println("--------------")
	add()
	fmt.Println("--------------")
	loadingImage()
	fmt.Println("--------------")
	GetHttp()

}
