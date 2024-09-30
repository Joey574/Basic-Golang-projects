package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	tpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contacthandler)
	http.HandleFunc("/login", loginHandler)

	// untested - idk which one takes priority
	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/", generalHandler)

	http.ListenAndServe(":8080", nil)
}

// security vulnerable with undeclared paths? idk. seems to be blank page with undeclared paths
func generalHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	tpl.ExecuteTemplate(w, path+".html", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}
