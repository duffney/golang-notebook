package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"strings"
)

type user struct {
	UserName string
	Password string //See sign-upPage for encrypted pwd example
}

var tpl *template.Template
var users = map[string]user{}
var sessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	users["test@email.com"] = user{"test@email.com", "P@ssw0rd"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.Handle("/favicon/ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", alreadyLoggedIn(req))
}

func login(w http.ResponseWriter, req *http.Request) {

	// already logged in?
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := users[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		if p != u.Password {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sessionUUID, _ := exec.Command("uuidgen").Output()
		cookie := &http.Cookie{
			Name:  "session",
			Value: strings.TrimSuffix(string(sessionUUID), "\n"),
		}
		http.SetCookie(w, cookie)
		sessions[cookie.Value] = un
		fmt.Println(sessions)
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.gohtml", users)
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
