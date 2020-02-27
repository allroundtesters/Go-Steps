package main

import (
	"fmt"
	"log"
	"net/http"
)

type Link string

func (s Link) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Structure struct {
	Greeting string
	Who      string
}

func (s Structure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func setupHandlers() {
	http.Handle("/string", Link("I'm a frayed knot."))
	http.Handle("/struct", &Structure{"Hello", "Gophers!"})
}

func main() {
	setupHandlers()
	log.Println("start server")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
