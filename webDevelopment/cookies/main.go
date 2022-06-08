package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/count", countCookie)
	http.HandleFunc("/expire", expireCookie)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "newCookie",
		Value: "a new value",
	})
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("newCookie")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, "Cookie value:", c)
}

func countCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("vists")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "vists",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}

func expireCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("vists")
	if err != nil {
		log.Println(err)
	}
	cookie.MaxAge = -1 // delete cookie
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
