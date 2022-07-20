package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

//to pass in a struct
func init() {
	tpl = template.Must(template.ParseFiles("templates/slice_struct/tpl.html"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Be cool, man",
	}

	sages := []sage{buddha, jesus}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
