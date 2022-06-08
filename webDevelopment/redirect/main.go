package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/posts", posts)
	http.HandleFunc("/images", images)
	http.HandleFunc("/img", img)
	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/thankyou", thankYou)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at home:", req.Method, "\n\n")
}

func blog(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at blog:", req.Method)
	// redirect with header
	w.Header().Set("Location", "/posts")
	w.WriteHeader(http.StatusSeeOther) //Status determines the type of redirect
}

func posts(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at posts:", req.Method)
	//io.WriteString(w, "HTTP request method:"+" "+req.Method)
	io.WriteString(w, "Redirected from /blog to /posts.")
}

func images(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at images:", req.Method)
	//clear browser cache to reset redirects
	http.Redirect(w, req, "/img", http.StatusMovedPermanently)
}

func img(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at img:", req.Method)
	io.WriteString(w, "Redirected from /images to /img.")
}

func subscribe(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		http.Redirect(w, req, "/thankyou", http.StatusSeeOther)
	}

	fmt.Println("Request method:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func thankYou(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request method at /thankyou:", req.Method)
	io.WriteString(w, "Thank you for subscribing!")
}
