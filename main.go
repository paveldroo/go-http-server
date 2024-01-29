package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "Here is index page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "Here is a dog page")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "Here is a me page. My name is Yoyo!")
}
