package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

var tpl *template.Template
var sessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandleFunc("/", setUuid)
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

	if _, ok := sessions[cookie.Value]; ok {
		http.Error(w, "Username already taken", http.StatusForbidden)
		return
	}

	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		sessions[cookie.Value] = userName
		fmt.Println(sessions)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	fmt.Println(cookie) //output this as a json object
}
