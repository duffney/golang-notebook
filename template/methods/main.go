package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type contact struct {
	Name   string
	Email  string
	GitHub string
}

func (c contact) NameToUpper() string {
	return strings.ToUpper(c.Name)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	c := contact{
		Name:   "Todd McLeod",
		Email:  "todd@email.com",
		GitHub: "@GoesToEleven",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", c)
	if err != nil {
		log.Fatalln(err)
	}
}
