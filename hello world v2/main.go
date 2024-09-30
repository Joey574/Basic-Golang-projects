package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/about", aboutHandle)
	http.HandleFunc("/", homeHandle)

	http.ListenAndServe(":8080", nil)
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.URL.Path:, %s", r.URL.Path)
}

func aboutHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our Gopher powered website.")
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello and Welcome!")
}
