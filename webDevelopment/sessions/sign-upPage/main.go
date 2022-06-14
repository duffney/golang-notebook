package main

import (
	"net/http"
	"text/template"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var users = map[string]user{}
var sessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func signup(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
