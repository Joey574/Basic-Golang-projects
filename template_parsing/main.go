package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	//tpl, _ = template.ParseFiles("index.html")
	//tpl, _ = template.ParseFiles("data1/index2.html")
	//tpl, _ = template.ParseFiles("data1/data2/index3.html")

	// move up one dir and grab index4
	//tpl, _ = template.ParseFiles("../index4.html")

	tpl, _ = tpl.ParseFiles("index.html")

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
