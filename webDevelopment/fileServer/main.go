package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/favicion", favicon)
	http.HandleFunc("/gopher", gopher)
	http.HandleFunc("/fromTemplate", fromTemplate)
	http.ListenAndServe(":8080", nil)
}

func gopher(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/golang-gopher.jpeg">`)
}

func fromTemplate(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func favicon(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "/favicon.ico")
}
