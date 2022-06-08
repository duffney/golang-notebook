package main

import (
	"html/template"
	"log"
	"os"
)

type subscriber struct {
	Name   string
	Email  string
	Source string
}

type emailList struct {
	subscribers []subscriber
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	subs := []subscriber{
		{
			Name:   "Abel",
			Email:  "abel@email.com",
			Source: "twitter",
		},
		{
			Name:   "Francis",
			Email:  "francis@email.com",
			Source: "reddit",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", subs)
	if err != nil {
		log.Fatalln(err)
	}
}
