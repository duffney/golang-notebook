package main

import (
	"fmt"
	"net/http"
)

type handler string

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello from the server!</h1>")
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
