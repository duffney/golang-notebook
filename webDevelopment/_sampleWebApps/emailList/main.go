package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var data subscriber

type subscriber struct {
	FirstName string
	LastName  string
	Email     string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/thanks", thanks)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func subscribe(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		e := req.FormValue("emailaddress")

		data = subscriber{f, l, e}
		http.Redirect(w, req, "/thanks", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "subscribe.gohtml", nil)
}

func thanks(w http.ResponseWriter, req *http.Request) {

	tpl.ExecuteTemplate(w, "thanks.gohtml", data)
	fmt.Println(data.FirstName)
}
