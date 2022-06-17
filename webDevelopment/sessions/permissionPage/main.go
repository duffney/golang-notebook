package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
)

type user struct {
	UserName string
	Password string
	Role     string
}

var tpl *template.Template
var users = map[string]user{}
var sessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	users["admin@email.com"] = user{"admin@email.com", "P@ssw0rd", "admin"}
	users["user@email.com"] = user{"user@email.com", "P@ssw0rd", "user"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/admin", admin)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
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
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func admin(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	fmt.Println(u)
	if u.Role != "admin" {
		http.Error(w, "You must be an admin to access this page", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "paywall.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("session")
	// delete the session
	delete(sessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sessionUUID, _ := exec.Command("uuidgen").Output()
		c = &http.Cookie{
			Name:  "session",
			Value: strings.TrimSuffix(string(sessionUUID), "\n"),
		}

	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := sessions[c.Value]; ok {
		u = users[un]
	}
	return u
}
