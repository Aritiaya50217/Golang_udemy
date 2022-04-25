package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type inventory map[string]float64

func (iv inventory) items(w http.ResponseWriter, r *http.Request) {
	for k, v := range iv {
		fmt.Fprintf(w, "%s : %v\n", k, v)
	}
}

func (iv inventory) price(w http.ResponseWriter, r *http.Request) {
	searchItem := r.URL.Query().Get("item")
	price, ok := iv[searchItem]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not item: %s", searchItem)
		return
	}
	fmt.Fprintf(w, "%v", price)
}

func (iv inventory) notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "sorry, no such page: %s", r.URL)
}

func main() {
	log.SetFlags(0)
	log.Println("pId : ", os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 3.99,
	}

	// http://localhost:8080/price?item=orange

	//mux := http.NewServeMux()
	//mux.Handle("/items", http.HandlerFunc(inven.items))
	//mux.Handle("/price", http.HandlerFunc(inven.price))
	//mux.Handle("/", http.HandlerFunc(inven.notFound))
	//
	//
	http.HandleFunc("/items", inven.items)
	http.HandleFunc("/price", inven.price)
	http.HandleFunc("/", inven.notFound)
	http.ListenAndServe(":8080", nil)
}
