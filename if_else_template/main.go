package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Language string
	Member   bool
}

var tpl *template.Template
var u User

func main() {
	//u = User{"Bob Smith", "English", false}
	//u = User{"Juan Hernandez", "Spanish", true}
	//u = User{"Zhang Wei", "Mandarin", true}
	u = User{"007", "?", true}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/welcome", welcomeHandler)

	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "member2.html", u)
}
