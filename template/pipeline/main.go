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
	"splitString":   splitToString,
}

var tpl *template.Template

func upperCase(s string) string {
	upperString := strings.ToUpper(s)
	return upperString
}

func splitToString(s string) string {
	split := strings.Split(s, "-")
	join := strings.Join(split, " ")
	return join
}

func init() {
	// the funcMap needs to be declared before the template is parsed
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("index.gohtml"))
}

func main() {
	// Call a function in a template
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "pipleline-template")
	if err != nil {
		log.Fatalln(err)
	}
}
