package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//type myHandler func(http.ResponseWriter, *http.Request)
//
//func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	m(w, r)
//}

type inventory map[string]float64

func (iv inventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Println("r.url", r.URL.Path)

	switch r.URL.Path {
	case "/items":
		for k, v := range iv {
			//w.Write([]byte("handle items"))
			fmt.Fprintf(w, "%s : %v\n", k, v)
		}
	//	http://localhost:8080/price?item=orange
	case "/price":
		searchItem := r.URL.Query().Get("item")
		price, ok := iv[searchItem]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "not item: %s", searchItem)
			return
		}
		fmt.Fprintf(w, "%v", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found page: %s", r.URL)
	}

	//w.Write([]byte("Hello"))
}

func main() {
	log.SetFlags(0)
	log.Println("pId : ", os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 3.99,
	}
	http.ListenAndServe(":8080", inven)
}
