package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", valueByURL)
	http.HandleFunc("/post", valueByBody)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//localhost:8080/?id=@joshduffney
//'id' is the identifier and @joshduffney is the value
func valueByURL(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at:", req.Method)

	v := req.FormValue("id")

	io.WriteString(w, "Value from URL is:"+v)
}

func valueByBody(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at /post:", req.Method)

	v := req.FormValue("id")

	tpl := template.Must(template.ParseGlob("*.gohtml"))

	err := tpl.ExecuteTemplate(w, "index.gohtml", v)
	if err != nil {
		log.Fatalln(err)
	}
}
