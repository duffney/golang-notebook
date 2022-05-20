package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	// Use 'ExecuteTemplate' to generate output to Stdout
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Use `ExecuteTemplate` to generate an HTML file
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	if err := tpl.ExecuteTemplate(newFile, "index.gohtml", nil); err != nil {
		log.Fatalln(err)
	}

	// Pass data to a template
	if err := tpl.ExecuteTemplate(newFile, "index.gohtml", "template"); err != nil { // "template" is the value
		log.Fatalln(err)
	}
}
