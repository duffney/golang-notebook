package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
)

type user struct {
	UserName string
	Password string
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
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")

		// does the user already exists?
		if _, ok := users[u]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create a session (map UUID to User)
		sessionUUID, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatalln(err)
		}
		cookie := &http.Cookie{
			Name:  "session",
			Value: strings.TrimSuffix(string(sessionUUID), "\n"),
		}
		http.SetCookie(w, cookie)
		sessions[cookie.Value] = u

		// add user to users database
		newUser := user{u, p}
		users[u] = newUser

		// redirect to index
		http.Redirect(w, req, "/", http.StatusSeeOther)
		fmt.Println(users)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)

}

func getUser(req *http.Request) user {
	var u user

	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	if un, ok := sessions[c.Value]; ok {
		u = users[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := sessions[c.Value]
	_, ok := users[un]
	return ok
}
