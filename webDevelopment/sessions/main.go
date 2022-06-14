package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	userName string
}

var tpl *template.Template
var sessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandleFunc("/", setUuid)
	http.HandleFunc("/signedIn", signedIn)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// store UUID in a cookie
func setUuid(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		uuidWithHyphen := uuid.New().String()
		cookie := &http.Cookie{
			Name:     "session",
			Value:    uuidWithHyphen,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)

	if userName, ok := sessions[cookie.Value]; ok {
		fmt.Println(userName)
		http.Redirect(w, req, "/signedIn", http.StatusSeeOther)
	}

	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		sessions[cookie.Value] = userName
		fmt.Println(sessions)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)

}

func signedIn(w http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("session")
	userName := sessions[cookie.Value]
	tpl.ExecuteTemplate(w, "loggedin.gohtml", userName)
}
