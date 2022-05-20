package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type user struct {
	name    string
	twitter string
	admin   bool
}

var funcMap = template.FuncMap{
	"makeUpperCase": upperCase,
}

var tpl *template.Template

func upperCase(s string) string {
	upperString := strings.ToUpper(s)
	return upperString
}

func init() {
	// the funcMap needs to be declared before the template is parsed
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("index.gohtml"))
}

func main() {
	// Call a function in a template
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "func-template")
	if err != nil {
		log.Fatalln(err)
	}
}
