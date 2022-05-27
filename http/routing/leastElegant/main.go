package main

import (
	"io"
	"net/http"
)

type handler string

func (m handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/about":
		io.WriteString(res, "About page")
	case "/blog":
		io.WriteString(res, "Blog page")
	default:
		io.WriteString(res, "Home page")
	}
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
