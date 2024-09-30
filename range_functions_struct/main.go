package main

import (
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

var Todo []task

var tpl *template.Template

func main() {

	//Todo = []task{{"give dog a bath", true}, {"mow the lawn", false}, {"pickup groceries", false}, {"take out trash", false}, {"paint kitchen", false}, {"feed dog", true}, {"water plants", false}}
	Todo = []task{{"give the dog a bath", true}, {"mow the lawn", true}, {"pickup groceries", true}, {"take out trash", true}, {"paint kitchen", false}, {"feed dog", true}, {"water plants", true}}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/todo", todoHandler)

	http.ListenAndServe(":8080", nil)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "todolist.html", Todo)
}
