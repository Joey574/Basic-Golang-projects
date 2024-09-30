package main

import (
	"html/template"
	"net/http"
)

type prodSpec struct {
	Size   string
	Weight float32
	Descr  string
}

type product struct {
	ProdID int
	Name   string
	Cost   float64
	Specs  prodSpec
}

var tpl *template.Template
var prod1 product

func main() {
	prod1 = product{
		ProdID: 15,
		Name:   "Really cool product",
		Cost:   899,
		Specs: prodSpec{
			Size:   "150 x 80 x 7 mm",
			Weight: 65,
			Descr:  "Coolest product",
		},
	}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/productinfo", productInfoHandler)
	http.ListenAndServe(":8080", nil)
}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "productinfo2.html", prod1)
}
